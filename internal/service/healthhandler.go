package service

import (
	"encoding/json"
	"log"
	"net/http"
)

// healthResponse is the json data sent upon a successful health req/resp sequence
type healthResponse struct {
	Health string `json:"health"`
}

// HealthHandler is responsible for checking the health of this service.
// Accessed through /{version}/health
func (s *Service) HealthHandler() http.HandlerFunc {
	responseBody := &healthResponse{Health: "healthy"}

	return func(resp http.ResponseWriter, req *http.Request) {
		if err := json.NewEncoder(resp).Encode(responseBody); err != nil {
			log.Println(err)
			resp.WriteHeader(http.StatusInternalServerError)
		}
	}
}
