package main

import (
	"net/http"
	"os"

	"puffin/cmd/server"
	"puffin/libs/database"
)

func main() {
	db := database.GetConnection(os.Getenv("DATABASE_DSN"))
	database.Migrate(db)

	http.ListenAndServe(
		":8008",
		server.CreateServer(
			&server.Options{
				DB: db,
			}),
	)
}
