package main

import (
	"net/http"

	"puffin/cmd/server"
)

func main() {
	http.ListenAndServe(":8008", server.CreateServer())
}
