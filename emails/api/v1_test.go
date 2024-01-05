package api_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"puffin/cmd/server"
	"puffin/emails"
	"puffin/internal/tests"

	"github.com/stretchr/testify/suite"
)

type EmailAPISuite struct {
	tests.SuiteWithDB
}

func TestEmailAPISuite(t *testing.T) {
	suite.Run(t, new(EmailAPISuite))
}

func (s *EmailAPISuite) TestCreateTemplate() {
	router := server.CreateServer(&server.Options{DB: s.DB})

	body := []byte(`{"name":"Test Name","content":"<h1>Test Content</h1>"}`)
	req := tests.CreateRequest("POST", "/api/v1/templates", body)
	response := tests.RecordCall(req, router)

	responseBody := emails.Template{}
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	s.Equal(http.StatusCreated, response.Code)
	s.Equal("Test Name", responseBody.Name)
	s.Equal("<h1>Test Content</h1>", responseBody.Content)
}

func (s *EmailAPISuite) TestGetTemplates() {
	router := server.CreateServer(&server.Options{DB: s.DB})

	// Create a template
	body := []byte(`{"name":"Test Name","content":"<h1>Test Content</h1>"}`)
	req := tests.CreateRequest("POST", "/api/v1/templates", body)
	tests.RecordCall(req, router)

	req = tests.CreateRequest("GET", "/api/v1/templates", nil)
	response := tests.RecordCall(req, router)

	responseBody := []emails.Template{}
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	s.Equal(http.StatusOK, response.Code)
	s.Equal(1, len(responseBody))
}
