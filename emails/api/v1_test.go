package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"puffin/cmd/server"
	"puffin/emails"

	"github.com/stretchr/testify/assert"
)

func record(req *http.Request, handler http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

func TestCreateTemplate(t *testing.T) {
	router := server.CreateServer()

	body := []byte(`{"name":"Test Name","content":"<h1>Test Content</h1>"}`)
	req, _ := http.NewRequest("POST", "/api/v1/templates", bytes.NewBuffer(body))
	response := record(req, router)

	responseBody := emails.Template{}
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, "Test Name", responseBody.Name)
	assert.Equal(t, "<h1>Test Content</h1>", responseBody.Content)
}

func TestGetTemplates(t *testing.T) {
	router := server.CreateServer()

	req, _ := http.NewRequest("GET", "/api/v1/templates", nil)
	response := record(req, router)

	responseBody := []emails.Template{}
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, 2, len(responseBody))
}
