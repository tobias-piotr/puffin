package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

// Decode decodes the request body into the target.
func Decode(r *http.Request, target any) error {
	return json.NewDecoder(r.Body).Decode(target)
}

// Validate validates a struct data.
func Validate(data any) error {
	return validate.Struct(data)
}

// DecodeAndValidate is a shortcut for Decode and Validate.
// It also parses the error into an APIError.
func DecodeAndValidate(r *http.Request, target any) error {
	if err := json.NewDecoder(r.Body).Decode(target); err != nil {
		return NewAPIError(http.StatusBadRequest, "Invalid request body", nil)
	}

	err := validate.Struct(target)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errors := make(map[string]string)
		// Not super happy about this
		for _, e := range validationErrors {
			errors[strings.ToLower(e.Field())] = e.Tag()
		}
		return NewAPIError(http.StatusUnprocessableEntity, "Unprocessable entity", errors)
	}

	return nil
}
