package common

import "os"

type Config struct {
}

// ConfigInstance .
var ConfigInstance = new(Config)

func (config *Config) Get(key string) string {
	value := os.Getenv(key)
	return value
}

func (config *Config) GetPortNumber() string {

	portNumber := config.Get("APPHOST_PortNumber")
	return portNumber
}

func (config *Config) GetStorageAccountName() string {

	value := config.Get("APPHOST_StorageAccountName")
	return value
}

func (config *Config) GetStorageAccountKey() string {

	value := config.Get("APPHOST_StorageAccountKey")
	return value
}

func (config *Config) GetTimeOutForService() string {

	value := config.Get("APPHOST_TimeOutForService")
	return value
}

func (config *Config) GetThresholdLengthForLongRunningQualification() string {

	value := config.Get("APPHOST_ThresholdLengthForLongRunningQualification")
	return value
}

func (config *Config) GetApplicationInsightKey() string {

	value := config.Get("APPHOST_ApplicationInsightKey")
	return value
}
