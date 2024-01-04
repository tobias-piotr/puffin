package api

import (
	"net/http"

	"puffin/emails"
	"puffin/emails/database"
	"puffin/libs/api"

	"github.com/go-chi/chi/v5"
)

type EmailHandler struct {
	service emails.EmailService
}

func (h *EmailHandler) CreateNewTemplate(w http.ResponseWriter, r *http.Request) {
	data := &emails.TemplateData{}
	if err := api.DecodeAndValidate(r, data); err != nil {
		api.RespondWithErr(w, err)
		return
	}

	t, err := h.service.CreateNewTemplate(data)
	if err != nil {
		api.RespondWithErr(w, err)
		return
	}

	api.Respond(w, http.StatusCreated, t)
}

func (h *EmailHandler) GetTemplates(w http.ResponseWriter, _ *http.Request) {
	t, err := h.service.GetTemplates()
	if err != nil {
		api.RespondWithErr(w, err)
		return
	}

	api.Respond(w, http.StatusOK, t)
}

func Register(r chi.Router) {
	handler := EmailHandler{emails.NewEmailService(database.EmailRepository{}, nil)}

	r.Route("/templates", func(r chi.Router) {
		r.Post("/", handler.CreateNewTemplate)
		r.Get("/", handler.GetTemplates)
	})
}
