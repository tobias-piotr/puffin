package database

import (
	"log/slog"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnection(dsn string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		slog.Error("Error connecting to database", "error", err.Error())
		panic(err)
	}
	return db
}
