package api

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	StatusCode int
	Err        error `json:"-"`
	Status     string
	Detail     string
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}
