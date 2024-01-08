package app

import (
	"net/http"

	"puffin/emails"
	"puffin/emails/database"
	"puffin/libs/smtp"
	"puffin/templates/pages"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type EmailAppHandler struct {
	service emails.EmailService
}

func (h *EmailAppHandler) GetTemplates(w http.ResponseWriter, r *http.Request) {
	templ.Handler(pages.ErrorPage()).ServeHTTP(w, r)
}

func Register(r chi.Router, db *sqlx.DB, smtpDialer smtp.Dialer) {
	service := emails.NewEmailService(
		database.NewEmailRepository(db),
		smtp.NewSmtpClient(smtpDialer),
	)
	appHandler := EmailAppHandler{service}

	r.Get("/", appHandler.GetTemplates)
}
