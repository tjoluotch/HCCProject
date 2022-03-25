package service

import "net/http"

// HealthHandler is responsible for checking the health of this service.
// Accessed through /{version}/health
func (s *Service) HealthHandler() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {

	}
}
