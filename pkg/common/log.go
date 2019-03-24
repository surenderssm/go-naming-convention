package common

import (
	"log"
	"strconv"
	"time"

	"github.com/Microsoft/ApplicationInsights-Go/appinsights"
)

// LogClient .
type LogClient struct {
	client appinsights.TelemetryClient
}

// Logger .
var Logger *LogClient

// NewLogger ...
func NewLogger(instrumentationKey string) *LogClient {
	logger := new(LogClient)

	if instrumentationKey == "" {
		logger.client = nil
	} else {
		telemetryConfig := appinsights.NewTelemetryConfiguration(instrumentationKey)

		// TODO : this is kept low to ensure even if there is low traffic data goes in real time

		telemetryConfig.MaxBatchSize = 5
		telemetryConfig.MaxBatchInterval = 5 * time.Second

		logger.client = appinsights.NewTelemetryClientFromConfig(telemetryConfig)
		// logger.client = appinsights.NewTelemetryClient(instrumentationKey)
	}
	return logger
}

// Info log message
func (logClient *LogClient) Info(message string) {

	if logClient != nil {
		logClient.client.TrackTrace(message, appinsights.Verbose)
	}
	log.Println(appinsights.Verbose, message)
}

// Error log error
func (logClient *LogClient) Error(message string) {

	if logClient != nil {
		logClient.client.TrackTrace(message, appinsights.Error)
	}
	log.Println(appinsights.Verbose, message)
}

// Request log a request
func (logClient *LogClient) Request(method string, operatonName string, duration time.Duration, statusCode int) {

	code := strconv.Itoa(statusCode)
	if logClient != nil {
		logClient.client.TrackRequest(method, operatonName, duration, code)
	}
	log.Println("Request", method, operatonName, duration, code)
}
