package service

import "net/http"

// VisitsHandler is responsible processing the bathroom visits file and returning bathroom visits data to the client.
// Accessed through /{version}/visits
func (s *Service) VisitsHandler() http.HandlerFunc {

	return func(resp http.ResponseWriter, req *http.Request) {

	}
}
