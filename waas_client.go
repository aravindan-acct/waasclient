package waasclient

import (
	"net/http"
	"time"
)

var Client *http.Client

// Barracuda waf-as-a-service API Base Endpoint
const URL = "https://api.waas.barracudanetworks.com/v2/waasapi/"

const (
	MaxIdleConnections int = 20
	RequestTimeout     int = 30
)

// Prepares a HTTP client to be used for making api calls to waf-as-a-service

func WaasClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}
	return client
}
