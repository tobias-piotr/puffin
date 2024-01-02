package server

import (
	"fmt"
	"net/http"
	"time"

	emails "puffin/emails/api"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func (h *HealthResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func CheckHealth(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, &HealthResponse{Status: "ok"})
}

func CreateServer(prefix string) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route(fmt.Sprintf("/%s/api", prefix), func(r chi.Router) {
		r.Get("/health", CheckHealth)

		r.Route("/v1", func(r chi.Router) {
			emails.Register(r)
		})
	})

	return r
}
