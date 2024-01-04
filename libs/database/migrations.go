package database

import (
	"embed"
	"log/slog"

	"github.com/jmoiron/sqlx"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var fs embed.FS

func Migrate(conn *sqlx.DB) {
	goose.SetBaseFS(fs)

	if err := goose.SetDialect("postgres"); err != nil {
		slog.Error("Couldn't set dialect", "error", err.Error())
		panic(err)
	}

	if err := goose.Up(conn.DB, "migrations"); err != nil {
		slog.Error("Migrating database failed", "error", err.Error())
		panic(err)
	}
}
