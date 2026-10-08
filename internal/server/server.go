package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PranavJoshi2893/frontend_mentor/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:              ":" + cfg.Port,
			Handler:           handler,
			ReadHeaderTimeout: 5 * time.Second,
			ReadTimeout:       15 * time.Second,
			WriteTimeout:      30 * time.Second,
			IdleTimeout:       60 * time.Second,
		},
	}
}

func (s *Server) Run() error {
	errCh := make(chan error, 1)

	go func() {
		slog.Info("http server is running at", "port", s.httpServer.Addr)
		err := s.httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errCh:
		return fmt.Errorf("failed to start http server: %w", err)

	case <-quit:
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		slog.Info("Shutting down http server...")
		err := s.httpServer.Shutdown(ctx)
		if err != nil {
			return fmt.Errorf("forced shutdown: %w", err)
		}
		slog.Info("http server gracefully shutdown")

		return nil
	}

}
