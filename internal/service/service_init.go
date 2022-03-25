package service

import (
	"fmt"
	"net/http"
	"time"
)

// newServer returns a *http.Server with the service router, host and port assigned
func (s *Service) newServer() *http.Server {
	return &http.Server{
		Handler:      s.router,
		Addr:         fmt.Sprintf("%s:%d", serviceHost, servicePort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
