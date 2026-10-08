package main

import (
	"log/slog"
	"os"

	"github.com/PranavJoshi2893/frontend_mentor/internal/config"
	"github.com/PranavJoshi2893/frontend_mentor/internal/router"
	"github.com/PranavJoshi2893/frontend_mentor/internal/server"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	r := router.Routes()
	srv := server.NewServer(cfg, r)

	if err := srv.Run(); err != nil {
		slog.Error("failed to start http server", "error", err)
		os.Exit(1)
	}

}
