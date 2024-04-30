package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/httprate"
	"github.com/mauriciofsnts/gofast/pkg/server/api/health"
)

func RouteApp(root *chi.Mux) {

	root.Route("/api/v1", func(r chi.Router) {
		// Default endpoints
		r.Get("/health", health.Health)

		// Rate limited endpoints
		// It applies a rate limit of 5 requests per IP address per minute to the "/api/v1/users" endpoint.
		// The rate limit is enforced using the httprate.LimitByIP middleware.
		r.Route("/users", func(r chi.Router) {
			// Apply rate limit of 5 requests per IP address per minute
			r.Use(httprate.LimitByIP(5, 1*time.Minute))
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("users"))
			})
		})
	})
}
