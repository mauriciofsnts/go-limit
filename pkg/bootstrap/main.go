package bootstrap

import (
	"log/slog"
	"os"

	"github.com/mauriciofsnts/gofast/pkg/config"
	"github.com/mauriciofsnts/gofast/pkg/ctx"
	"github.com/mauriciofsnts/gofast/pkg/server"
)

func Start(cfg *config.Config) {
	setupLog(cfg)

	slog.Info("Starting application")
	slog.Debug("If you see this message, the log level is set to debug")

	services := &ctx.Services{
		Config: cfg,
	}

	err := server.StartServer(services)

	if err != nil {
		slog.Error("Failed to start server:", "err", err)
		os.Exit(1)
	}
}
