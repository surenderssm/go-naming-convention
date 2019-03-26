package common

import (
	"os"
	"strconv"
)

type Config struct {
}

// ConfigInstance .
var ConfigInstance = new(Config)

func (config *Config) Get(key string) string {
	value := os.Getenv(key)
	return value
}

func (config *Config) GetPortNumber() string {

	portNumber := config.Get("PortNumber")
	return ":" + portNumber
}

func (config *Config) GetStorageAccountName() string {

	value := config.Get("StorageAccountName")
	return value
}

func (config *Config) GetStorageAccountKey() string {

	value := config.Get("StorageAccountKey")
	return value
}

func (config *Config) GetOxfordTimeOutForService() int {

	value := config.Get("OxfordTimeOutForService")
	i, _ := strconv.Atoi(value)
	return i
}

func (config *Config) GetApplicationInsightKey() string {

	value := config.Get("ApplicationInsightKey")
	return value
}

func (config *Config) GetOxfordBaseURL() string {

	value := config.Get("OxfordBaseURL")
	return value
}

func (config *Config) GetOxfordAppId() string {

	value := config.Get("OxfordAppId")
	return value
}

func (config *Config) GetOxfordAppSecret() string {

	value := config.Get("OxfordAppSecret")
	return value
}

func (config *Config) GetContainerName() string {

	value := config.Get("ContainerName")
	return value
}

func (config *Config) GetWordFileName() string {
	value := config.Get("WordFileName")
	return value
}
