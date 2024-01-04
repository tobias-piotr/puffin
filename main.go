package main

import (
	"net/http"

	"puffin/cmd/server"
)

func main() {
	// TODO: Add dependecies here
	http.ListenAndServe(":8008", server.CreateServer())
}
