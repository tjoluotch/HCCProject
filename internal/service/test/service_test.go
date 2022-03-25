package service_test

import (
	"HCCProject/internal/service"
	"encoding/json"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("server tests", func() {
	var (
		healthResponse service.HealthResponse
		expHealthResp  service.HealthResponse

		visitResponse service.VisitResponse

		serviceObj     *service.Service
		request        *http.Request
		handler        http.Handler
		responseWriter *httptest.ResponseRecorder
	)

	Describe("handle request to {version}/health", func() {
		BeforeEach(func() {
			expHealthResp = service.HealthResponse{Health: "healthy"}

			serviceObj = service.NewService()
			handler = serviceObj.HealthHandler()
			request = httptest.NewRequest(http.MethodGet, service.HealthURI, nil)
			responseWriter = httptest.NewRecorder()
		})

		It("processes the {version}/health request successfully", func() {
			handler.ServeHTTP(responseWriter, request)
			respData := responseWriter.Body.Bytes()
			err := json.Unmarshal(respData, &healthResponse)

			Expect(responseWriter.Code).To(Equal(http.StatusOK))
			Expect(err).To(BeNil())
			Expect(healthResponse.Health).To(Equal(expHealthResp.Health))
		})
	})

	Describe("handle request to {version}/visits", func() {
		BeforeEach(func() {
			serviceObj = service.NewService()
			handler = serviceObj.VisitsHandler()
			request = httptest.NewRequest(http.MethodGet, service.VisitsURI, nil)
			responseWriter = httptest.NewRecorder()
		})

		It("processes the {version}/visits request successfully", func() {
			handler.ServeHTTP(responseWriter, request)
			err := json.Unmarshal(responseWriter.Body.Bytes(), &visitResponse)

			Expect(responseWriter.Code).To(Equal(http.StatusOK))
			Expect(err).To(BeNil())
			Expect(visitResponse).To(MatchAllFields(Fields{
				"Data": Not(BeEmpty()),
			}))
		})
	})

})
