package service

import "HCCProject/internal/filehandler"

// addFileHandler sets the file handler field within the service object to a new filehandler object
func (s *Service) addFileHandler() {
	s.fh = filehandler.NewFileHandler()
}
