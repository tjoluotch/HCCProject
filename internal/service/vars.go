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
	VisitsURI = fmt.Sprintf("%s%s", version, visitsEndpoint)
	HealthURI = fmt.Sprintf("%s%s", version, healthEndpoint)
)
