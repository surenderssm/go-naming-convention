package handlers

import (
	"encoding/json"
	"go-naming-convention/pkg/common"
	"go-naming-convention/pkg/processor"
	"go-naming-convention/pkg/repository"
	"net/http"
	"time"

	"github.com/rs/xid"
)

const tokenKey string = "token"
const caseTypeKey string = "type"
const trackingKey string = "trackingid"
const correlationKey string = "x-ms-correlation"

// NamingModel case model
type NamingModel struct {
	Token       string `json:"token"`
	Result      string `json:"result"`
	CaseType    string `json:"type"`
	TrackingID  string `json:"trackingID"`
	TrackingURL string `json:"trackingUrl"`
	Message     string `json:"message"`
	StatusCode  int    `json:"statusCode"`
}

// BlobStore for the system
var BlobStore *repository.BlobStore

// TokenProcessor processor for the system
var TokenProcessor *processor.Tokenizer

// threshold for token length allowed in GET
// TODO : can be moved to config
var thresholdLengthForTokenForSyncProcessing = 200
var thresholdLengthForTokenForAsyncProcessing = 1000

// Name process the tokens as per the given case
// TODO : have a mechanism to avoid throttling (429) by keeping a check on node signals / total ongoing request
func Name(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	token, caseType := parseNameRequest(r)
	var model *NamingModel

	if ok, itemModel := validateWork(token, caseType); !ok {
		model = itemModel
		// 400
		model.StatusCode = http.StatusBadRequest
	} else if (r.Method == "GET" && len(token) > thresholdLengthForTokenForSyncProcessing) || (r.Method == "POST" && len(token) > thresholdLengthForTokenForAsyncProcessing) {
		model := new(NamingModel)
		model.Token = token
		model.CaseType = caseType
		// 413
		model.StatusCode = http.StatusRequestEntityTooLarge
	} else if r.Method == "GET" {
		model = processName(token, caseType)
		// 200
		model.StatusCode = http.StatusOK
	} else if r.Method == "POST" {
		model = processNameLongRunning(token, caseType)
		// 202 - long running operation
		// caller has to track result at TrackingURL
		model.StatusCode = http.StatusAccepted
	}
	// Log the Request / info
	// TODO : log the custom event to keep track of varios aspect of the endpoint
	defer logRequestWithModel(r, model, time.Since(now))
	w.WriteHeader(model.StatusCode)
	json.NewEncoder(w).Encode(model)

}

// TrackWork .
func TrackWork(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	trackingID := getQuery(trackingKey, r)

	if trackingID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// HEAD instead of get can be perfomed to check the presence instead of failing in case of missing
	data, error := BlobStore.GetBlob(trackingID)
	model := NamingModel{}
	if error == nil {
		// blob is being returned
		// notion of TTL for blob should be handle to delete the allready served model
		// TTL can be kept as 24 hour or so to stop ever growing data
		json.Unmarshal([]byte(data), &model)
		model.StatusCode = http.StatusOK
	} else {
		// 202 if the request is still under process and results are not returned yet
		model = NamingModel{TrackingID: trackingID, TrackingURL: getTrackingURI(trackingID), StatusCode: http.StatusAccepted}
	}
	defer logRequestWithModel(r, &model, time.Since(now))
	w.WriteHeader(model.StatusCode)
	json.NewEncoder(w).Encode(model)
}

func getTrackingURI(trackingID string) string {
	// TODO : concatenate with hostname to get full URI
	return "v1/track?" + trackingKey + "=" + trackingID
}

func processName(token string, caseType string) *NamingModel {
	result := TokenProcessor.ProcessToken(token, processor.NamingConventionCase(caseType))
	item := NamingModel{Token: token, Result: result, CaseType: caseType}
	return &item
}

func processNameLongRunning(token string, caseType string) *NamingModel {
	// generate a new trackingID
	trackingID := xid.New().String()
	// TODO : have a notion of worker pool with a threshold, as simply using go might exhaust all the resources
	// Like if threshols is reached 429 can be returned back
	// notion of producer/Consumer can also solve this proble
	// for example light weight Azure function can be poll on the queue to process the token and keep the reslt in the store
	go processTokensLongRunning(token, caseType, trackingID)
	item := NamingModel{Token: token, TrackingID: trackingID, CaseType: caseType, TrackingURL: getTrackingURI(trackingID)}
	return &item
}

func processTokensLongRunning(token string, caseType string, trackingID string) {
	result := TokenProcessor.ProcessToken(token, processor.NamingConventionCase(caseType))
	item := NamingModel{Token: token, Result: result, TrackingID: trackingID, CaseType: caseType}
	byteData, _ := json.Marshal(&item)
	data := string(byteData)
	BlobStore.CreateBlockBlob(trackingID, data)
}

func validateWork(token string, caseType string) (bool, *NamingModel) {
	var isValid bool
	// check for valid token
	isValid = processor.IsValidToken(token)

	model := NamingModel{Token: token, CaseType: caseType}
	if !isValid {
		model.Message = "InvalidToken"
		return isValid, &model
	}

	isValid = processor.IsValidCaseType(caseType)

	if !isValid {
		model.Message = "InvalidCase"
		return isValid, &model
	}
	return true, nil
}

func logRequest(r *http.Request, duration time.Duration, statusCode int) {
	common.Logger.Request(r.Method, r.URL.String(), duration, statusCode)
}

func logRequestWithModel(r *http.Request, model *NamingModel, duration time.Duration) {

	text, _ := json.Marshal(model)
	common.Logger.Info(string(text))
	logRequest(r, duration, model.StatusCode)
}

func getQuery(key string, r *http.Request) string {
	if keys, ok := r.URL.Query()[key]; ok {
		return keys[0]
	}
	return ""
}

// parseNameRequest get the token, case from teh given request
func parseNameRequest(r *http.Request) (string, string) {
	var token, caseType string

	if r.Method == "GET" {
		token = getQuery(tokenKey, r)
		caseType = getQuery(caseTypeKey, r)
	} else if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var model NamingModel
		err := decoder.Decode(&model)

		if err == nil {
			token = model.Token
			caseType = model.CaseType
		} else {
			// incase of any failure let token/case "",we are failed to extract it caller should handle with bad request or something
			common.Logger.Error(err.Error())
		}
		return token, caseType
	}

	return token, caseType
}
