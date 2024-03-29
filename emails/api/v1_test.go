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
	Dialer *tests.DummyDialer
}

func (s *EmailAPISuite) SetupSuite() {
	s.SuiteWithDB.SetupSuite()
	s.Dialer = &tests.DummyDialer{}
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

func (s *EmailAPISuite) TestCreateDuplicateTemplate() {
	router := server.CreateServer(&server.Options{DB: s.DB})

	// Create a template
	body := []byte(`{"name":"Test Name","content":"<h1>Test Content</h1>"}`)
	req := tests.CreateRequest("POST", "/api/v1/templates", body)
	tests.RecordCall(req, router)

	// Make the same request again
	req = tests.CreateRequest("POST", "/api/v1/templates", body)
	response := tests.RecordCall(req, router)

	s.Equal(http.StatusConflict, response.Code)
	s.Equal(`{"error":"Template with name Test Name already exists"}`, response.Body.String())
}

func (s *EmailAPISuite) TestGetTemplates() {
	router := server.CreateServer(&server.Options{DB: s.DB})

	// Create a template
	body := []byte(`{"name":"Test Name","content":"<h1>Test Content</h1>"}`)
	req := tests.CreateRequest("POST", "/api/v1/templates", body)
	tests.RecordCall(req, router)

	// Get all templates
	req = tests.CreateRequest("GET", "/api/v1/templates", nil)
	response := tests.RecordCall(req, router)

	responseBody := []emails.Template{}
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	s.Equal(http.StatusOK, response.Code)
	s.Equal(1, len(responseBody))
}

func (s *EmailAPISuite) TestSendEmail() {
	router := server.CreateServer(&server.Options{DB: s.DB, SmtpDialer: s.Dialer})

	// Create a template
	body := []byte(`{"name":"test","content":"<h1>Hello {{.name}}</h1>"}`)
	req := tests.CreateRequest("POST", "/api/v1/templates", body)
	tests.RecordCall(req, router)

	// Send an email
	body = []byte(`{"recipients":["test@gmail.com"],"template_name":"test","subject":"Test","data":{"name":"Test Name"}}`)
	req = tests.CreateRequest("POST", "/api/v1/emails", body)
	response := tests.RecordCall(req, router)

	s.Equal(http.StatusOK, response.Code)

	// Check that the email was saved
	req = tests.CreateRequest("GET", "/api/v1/emails", nil)
	response = tests.RecordCall(req, router)

	responseBody := []emails.Email{}
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	s.Equal(http.StatusOK, response.Code)
	s.Equal(1, len(responseBody))
}

func (s *EmailAPISuite) TestSendEmailNoTemplate() {
	router := server.CreateServer(&server.Options{DB: s.DB})

	body := []byte(`{"recipients":["test@gmail.com"],"template_name":"test","subject":"Test","data":{"name":"Test Name"}}`)
	req := tests.CreateRequest("POST", "/api/v1/emails", body)
	response := tests.RecordCall(req, router)

	s.Equal(http.StatusNotFound, response.Code)
	s.Equal(`{"error":"Template with name test does not exist"}`, response.Body.String())
}
