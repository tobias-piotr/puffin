package main

import (
	"net/http"
	"os"

	"puffin/cmd/server"
	"puffin/libs/database"

	"github.com/go-mail/mail"
)

// @title Puffin
// @version 0.1.0
// @description Sending emails made complicated
// @BasePath /puffin
func main() {
	db := database.GetConnection(os.Getenv("DATABASE_DSN"))
	database.Migrate(db)

	smptDialer := mail.NewDialer(os.Getenv("SMTP_HOST"), 587, os.Getenv("SMTP_EMAIL"), os.Getenv("SMTP_PASSWORD"))

	http.ListenAndServe(
		":8008",
		server.CreateServer(
			&server.Options{
				DB:         db,
				SmtpDialer: smptDialer,
			}),
	)
}
