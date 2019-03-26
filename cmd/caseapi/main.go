package main

import (
	"go-naming-convention/pkg/common"
	"go-naming-convention/pkg/handlers"
	"go-naming-convention/pkg/processor"
	"go-naming-convention/pkg/repository"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {

	log.Println("caseapiStarting" + time.Now().Format(time.RFC850))

	initializePlatform()
	log.Println("caseapiInitialized" + time.Now().Format(time.RFC850))
	handleRequest()
}

func handleRequest() {

	// http.HandleFunc("/v1/echo", handlers.Echo)
	http.HandleFunc("/v1/name", handlers.Name)
	http.HandleFunc("/v1/track", handlers.TrackWork)
	http.HandleFunc("/v1/ping", handlers.Ping)
	http.HandleFunc("/v1/health/ping", handlers.Ping)
	http.HandleFunc("/v1/health", handlers.Health)

	log.Println("caseapiListenAndServing" + time.Now().Format(time.RFC850))

	portNumber := common.ConfigInstance.GetPortNumber()
	// TODO: get port number from OS Environment
	// port := os.Getenv("PORT")
	// default handler is DefaultServeMux
	log.Fatal(http.ListenAndServe(portNumber, nil))
}

// Init prepare the handlers to handle upcoming request
//
func initializePlatform() {
	common.Logger = getLogger()
	store := getBlobStore()
	handlers.BlobStore = store
	handlers.TokenProcessor = getTokenProcessor(store)
}

func getTokenProcessor(blobStore *repository.BlobStore) *processor.Tokenizer {

	provider := getWordProvider()
	data := getOfflineWords(blobStore)
	repo := repository.NewWordRepository(data, provider)
	tokenizer := processor.NewTokenizer(repo)
	return tokenizer
}

func getLogger() *common.LogClient {

	key := common.ConfigInstance.GetApplicationInsightKey()
	logger := common.NewLogger(key)
	return logger
}

func getBlobStore() *repository.BlobStore {

	accountName := common.ConfigInstance.GetStorageAccountName()
	accountKey := common.ConfigInstance.GetStorageAccountKey()
	containerName := common.ConfigInstance.GetContainerName()
	store := repository.NewBlobStore(accountName, accountKey, containerName)
	return store
}

func getWordProvider() *repository.OxfordWordProvider {
	baseurl := common.ConfigInstance.GetOxfordBaseURL()
	appid := common.ConfigInstance.GetOxfordAppId()
	appSecret := common.ConfigInstance.GetOxfordAppSecret()
	threshold := common.ConfigInstance.GetOxfordTimeOutForService()
	wordProvider := repository.NewOxfordWordProvider(baseurl, appid, appSecret, threshold)
	return wordProvider
}

func getOfflineWords(store *repository.BlobStore) []string {

	fileName := common.ConfigInstance.GetWordFileName()
	data, error := store.GetBlob(fileName)

	if error != nil {
		common.Logger.Error("getOfflineWordsFailed-" + error.Error())
		return make([]string, 0)
	}
	items := strings.Split(data, "\n")
	common.Logger.Info("getOfflineWords-total-" + strconv.Itoa(len(items)))
	return items
}
