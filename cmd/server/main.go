package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/swaggo/http-swagger/v2"

	emails "puffin/emails/api"
	"puffin/libs/api"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "puffin/docs"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func CheckHealth(w http.ResponseWriter, _ *http.Request) {
	api.Respond(w, http.StatusOK, &HealthResponse{Status: "ok"})
}

func CreateServer(opts *Options) chi.Router {
	slog.Info("Creating server")

	prefix := os.Getenv("API_PREFIX")

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(NewLoggingMiddleware())
	r.Use(SetContentType)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route(fmt.Sprintf("%s/api", prefix), func(r chi.Router) {
		r.Get("/health", CheckHealth)

		r.Get("/docs/*", httpSwagger.Handler(
			httpSwagger.URL(fmt.Sprintf("%s/api/docs/doc.json", prefix)),
		))

		r.Route("/v1", func(r chi.Router) {
			emails.Register(r, opts.DB, opts.SmtpDialer)
		})
	})

	return r
}
