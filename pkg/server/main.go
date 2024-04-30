package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mauriciofsnts/gofast/pkg/ctx"
	"github.com/mauriciofsnts/gofast/pkg/server/router"
)

func StartServer(services *ctx.Services) error {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(servicesContext(services))

	router.RouteApp(r)

	bindAddr := fmt.Sprintf(":%d", services.Config.Http.Port)
	slog.Info("Starting server on %s", bindAddr)

	server := &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      r,
		Addr:         bindAddr,
	}

	return server.ListenAndServe()
}
