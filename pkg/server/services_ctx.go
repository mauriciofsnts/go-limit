package server

import (
	"context"
	"net/http"

	"github.com/mauriciofsnts/gofast/pkg/ctx"
)

func servicesContext(services *ctx.Services) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(
				w,
				r.WithContext(
					context.WithValue(r.Context(), ctx.ServicesKey, services),
				),
			)
		})
	}
}
