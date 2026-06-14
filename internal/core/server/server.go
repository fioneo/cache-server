package core_server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"
)

type HTTPServer struct {
	mux *http.ServeMux
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		mux: http.NewServeMux(),
	}
}

func (s *HTTPServer) Start(ctx context.Context) error {
	server := http.Server{
		Handler: s.mux,
		Addr:    ":8080",
	}

	errCh := make(chan error, 1)

	go func() {
		defer close(errCh)
		slog.Info("starting HTTP server", "addr", server.Addr)

		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		slog.Info("shutting down HTTP server")

		shutDownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(shutDownCtx); err != nil {
			_ = server.Close()

			slog.Error("http server shutdown failed", "error", err)
			return err
		}
		slog.Info("http server shutdown successful")
	case err := <-errCh:
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *HTTPServer) RegisterRoutes(routes ...*ApiRouter) {
	for _, router := range routes {
		s.mux.Handle("/", router)
	}
}
