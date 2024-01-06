package server

import (
	"puffin/libs/smtp"

	"github.com/jmoiron/sqlx"
)

type Options struct {
	DB         *sqlx.DB
	SmtpDialer smtp.Dialer
}
