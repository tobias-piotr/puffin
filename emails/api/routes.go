package api

import (
	"puffin/emails"
	"puffin/emails/database"
	"puffin/libs/smtp"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func Register(r chi.Router, db *sqlx.DB, smtpDialer smtp.Dialer) {
	service := emails.NewEmailService(
		database.NewEmailRepository(db),
		smtp.NewSmtpClient(smtpDialer),
	)
	v1Handler := EmailHandler{service}

	r.Route("/v1", func(r chi.Router) {
		r.Route("/templates", func(r chi.Router) {
			r.Post("/", v1Handler.CreateNewTemplate)
			r.Get("/", v1Handler.GetTemplates)
		})

		r.Route("/emails", func(r chi.Router) {
			r.Post("/", v1Handler.SendEmail)
			r.Get("/", v1Handler.GetEmails)
		})
	})

	// TODO: Add routes for htmx
}
