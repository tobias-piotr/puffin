package api

import (
	"net/http"

	"puffin/emails"
	"puffin/libs/api"

	"github.com/go-chi/chi/v5"
)

func CreateNewTemplate(w http.ResponseWriter, r *http.Request) {
	data := &emails.TemplateData{}
	if err := api.DecodeAndValidate(r, data); err != nil {
		api.RespondWithErr(w, err)
		return
	}

	srv := emails.EmailService{}

	t, err := srv.CreateNewTemplate()
	if err != nil {
		api.RespondWithErr(w, err)
		return
	}

	api.Respond(w, http.StatusCreated, t)
}

func Register(r chi.Router) {
	r.Route("/templates", func(r chi.Router) {
		r.Post("/", CreateNewTemplate)
	})
}
