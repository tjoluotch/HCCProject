package service

import (
	"HCCProject/internal/filehandler"
	"github.com/gorilla/mux"
)

// Service holds all the dependencies needed to start and execute requests on the server
type Service struct {
	fh     filehandler.FileHandler
	router *mux.Router
}

// NewService returns an empty *Service
func NewService() *Service {
	return &Service{}
}
