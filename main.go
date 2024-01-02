package main

import (
	"net/http"
	"os"

	"puffin/cmd/server"
)

func main() {
	prefix := os.Getenv("API_PREFIX")
	http.ListenAndServe(":8008", server.CreateServer(prefix))
}
