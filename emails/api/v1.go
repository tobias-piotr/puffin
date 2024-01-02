package api

import (
	"net/http"

	"puffin/emails"
	"puffin/libs/api"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type TemplateRequest struct {
	emails.TemplateData
}

func (t *TemplateRequest) Bind(r *http.Request) error {
	return nil
}

type TemplateResponse struct {
	emails.Template
}

func (t TemplateResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func CreateNewTemplate(w http.ResponseWriter, r *http.Request) {
	data := &TemplateRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, &api.ErrResponse{StatusCode: http.StatusBadRequest, Err: err, Status: "Bad Request", Detail: err.Error()})
		return
	}

	t := emails.Template{Name: "test", Content: "test"}
	render.Render(w, r, TemplateResponse{t})
}

func Register(r chi.Router) {
	r.Route("/templates", func(r chi.Router) {
		r.Post("/", CreateNewTemplate)
	})
}
