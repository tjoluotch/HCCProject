package service

import "fmt"

const (
	version        = "/v1"
	healthEndpoint = "/health"
	visitsEndpoint = "/visits"

	serviceHost = "127.0.0.1"
	servicePort = 3000
)

var (
	visitsURI = fmt.Sprintf("%s%s", version, visitsEndpoint)
	healthURI = fmt.Sprintf("%s%s", version, healthEndpoint)
)
