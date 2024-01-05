package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
)

// CreateRequest is a helper function to create a request with a body.
func CreateRequest(method string, url string, body []byte) *http.Request {
	url = os.Getenv("API_PREFIX") + url
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(body))
	return req
}

// RecordCall is a helper function to record the response of a handler.
func RecordCall(req *http.Request, handler http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}
