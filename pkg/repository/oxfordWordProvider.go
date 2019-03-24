package repository

import (
	"errors"
	"go-naming-convention/pkg/common"
	"net/http"
	"time"
)

// OxfordWordProvider word provider using Oxford api
// TODO : kind of an interface to accomodate multiple word provider for eg (oxford,wiki,...)
type OxfordWordProvider struct {
	BaseURL   string
	AppID     string
	AppSecret string
	NetClient *http.Client
	// can be use to keep check on throttling of provider,
	//if threshold is crossed smart decision like switching to other provider and all can be made
	TotalHit int
}

// NewOxfordWordProvider  provider calling oxford
func NewOxfordWordProvider(BaseURL string, appId string, appSecret string, timeoutThreshold int) *OxfordWordProvider {

	// TODO : check for valid arguments
	provider := new(OxfordWordProvider)
	provider.BaseURL = BaseURL     // "https://od-api.oxforddictionaries.com/api/v1/entries/en/
	provider.AppID = appId         //b7f2b4d2
	provider.AppSecret = appSecret // 27e47a2c635fba71e204b861f5e957b2
	provider.NetClient = new(http.Client)
	provider.NetClient.Timeout = time.Duration(timeoutThreshold * int(time.Millisecond))
	return provider
}

// WordExist check if the word exist
func (provider *OxfordWordProvider) WordExist(word string) (bool, error) {

	url := provider.BaseURL + word
	// For a word as simple as "ok" 20 KB of body is being returned in GET, whereas HEAD is ~ 30 bytes
	// HEAD instead of GET saves the data flow - cost and save memory,time,cpu cycles
	req, requestError := http.NewRequest("HEAD", url, nil)
	req.Header.Set("app_id", provider.AppID)
	req.Header.Set("app_key", provider.AppSecret)

	response, requestError := provider.NetClient.Do(req)
	provider.TotalHit = provider.TotalHit + 1

	if response != nil && response.StatusCode == 404 {
		return false, nil
	} else if response != nil && response.StatusCode == 200 {
		return true, nil
	}

	// either service has timed out or service is not available or some other issues
	// log error

	if requestError != nil {
		common.Logger.Error("OxfordWordProviderFailedWithError:" + requestError.Error())
	}

	return false, errors.New("ProviderNotServiceble")
}
