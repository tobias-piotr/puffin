package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// APIError represents an error that can be returned by the API.
type APIError struct {
	StatusCode int               `json:"-"`
	Status     string            `json:"error"`
	Detail     map[string]string `json:"detail,omitempty"`
}

func NewAPIError(statusCode int, status string, detail map[string]string) APIError {
	return APIError{statusCode, status, detail}
}

func (e APIError) Error() string {
	return e.Status
}

// RespondWithErr writes an APIError to the response.
func RespondWithErr(w http.ResponseWriter, err error) error {
	aerr, ok := err.(APIError)
	if !ok {
		slog.Error("Internal server error", "error", err)
		aerr = NewAPIError(http.StatusInternalServerError, "Internal server error", nil)
	}
	w.WriteHeader(aerr.StatusCode)
	res, err := json.Marshal(aerr)
	if err != nil {
		return err
	}
	w.Write(res)
	return nil
}

// Respond writes a marshalled data to the response.
func Respond(w http.ResponseWriter, status int, data any) error {
	w.WriteHeader(status)
	res, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.Write(res)
	return nil
}
