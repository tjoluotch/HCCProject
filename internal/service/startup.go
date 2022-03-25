package service

import "log"

// Start is responsible for service startup
func Start() {
	// create new service
	s := NewService()
	// initialise file handler dependency to the service
	s.addFileHandler()
	// initialise service endpoints
	s.routes()
	// start server
	log.Fatal(s.newServer().ListenAndServe())
}
