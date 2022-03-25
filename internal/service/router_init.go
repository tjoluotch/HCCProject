package service

import (
	"github.com/gorilla/mux"
	"net/http"
)

// routes is responsible for setting up the service routes and updating the router field within the service
func (s *Service) routes() {
	router := mux.NewRouter()
	router.Handle(visitsURI, s.VisitsHandler()).Methods(http.MethodGet)
	router.Handle(healthURI, s.HealthHandler()).Methods(http.MethodGet)
	s.router = router
}
