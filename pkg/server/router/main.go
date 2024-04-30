package router

import (
	"github.com/go-chi/chi"
	"github.com/mauriciofsnts/gofast/pkg/server/api/health"
)

func RouteApp(root *chi.Mux) {
	root.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", health.Health)
	})
}
