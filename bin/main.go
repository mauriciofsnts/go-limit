package main

import (
	"log/slog"
	"os"

	"github.com/mauriciofsnts/gofast/pkg/bootstrap"
	"github.com/mauriciofsnts/gofast/pkg/config"
)

const (
	DefaultConfigPath = "config.yaml"
)

func main() {
	cfg, err := config.LoadConfigFromFile(DefaultConfigPath)
	if err != nil {
		slog.Error("Failed to load config file:", "err", err)
		os.Exit(1)
	}

	bootstrap.Start(cfg)
}
