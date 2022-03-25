package service

import (
	"HCCProject/internal/filehandler"
	"encoding/json"
	"log"
	"net/http"
)

// visitResponse contains the json data for a successful responses regarding client visualisation data points
type visitResponse struct {
	Data []filehandler.Data `json:"data"`
}

// errorResponse contains error if we failed to calculate trend data
type errorResponse struct {
	Error string `json:"error"`
}

// VisitsHandler is responsible processing the bathroom visits file and returning bathroom visits data to the client.
// Accessed through /{version}/visits
func (s *Service) VisitsHandler() http.HandlerFunc {
	visitData, err := s.fh.CalculateTrendDataPoints()
	var dataPoints []filehandler.Data
	response := visitResponse{}

	return func(resp http.ResponseWriter, req *http.Request) {
		// if we failed to calculate trend data return 500 response and errorResponse type
		if err != nil {
			log.Println(err)
			resp.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(resp).Encode(&errorResponse{Error: err.Error()})
		}

		// populate data for response object
		for _, data := range visitData {
			dataPoints = append(dataPoints, *data)
		}
		// assign the slice to response object
		response.Data = dataPoints

		_ = json.NewEncoder(resp).Encode(&response)
	}
}
