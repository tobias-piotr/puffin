package api

import (
	"net/http"

	"puffin/emails"
	"puffin/emails/database"
	"puffin/libs/api"
	"puffin/libs/smtp"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
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

func (h *EmailHandler) SendEmail(w http.ResponseWriter, r *http.Request) {
	data := &emails.EmailData{}
	if err := api.DecodeAndValidate(r, data); err != nil {
		api.RespondWithErr(w, err)
		return
	}

	if err := h.service.SendEmail(data); err != nil {
		api.RespondWithErr(w, err)
		return
	}

	api.Respond(w, http.StatusOK, nil)
}

func Register(r chi.Router, db *sqlx.DB, smtpDialer smtp.Dialer) {
	handler := EmailHandler{
		emails.NewEmailService(
			database.NewEmailRepository(db),
			smtp.NewSmtpClient(smtpDialer),
		),
	}

	r.Route("/templates", func(r chi.Router) {
		r.Post("/", handler.CreateNewTemplate)
		r.Get("/", handler.GetTemplates)
	})

	r.Route("/emails", func(r chi.Router) {
		r.Post("/", handler.SendEmail)
	})
}
