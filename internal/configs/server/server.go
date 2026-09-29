package app_server

import (
	"authService/internal/midelware"
	"context"
	"errors"
	"fmt"
	"net/http"
)

type httpServer struct {
	config     ServerConfig
	ServerMux  *http.ServeMux
	Middleware []midelware.Middleware
}

func NewServer(config ServerConfig, middleware ...midelware.Middleware) *httpServer {
	return &httpServer{
		config:     config,
		ServerMux:  http.NewServeMux(),
		Middleware: middleware,
	}
}

func (s *httpServer) Run(ctx context.Context) error {

	mux := midelware.MiddlewareChain(s.ServerMux, s.Middleware...)

	server := &http.Server{
		Addr:    s.config.addr,
		Handler: mux,
	}

	ch := make(chan error, 1)

	go func() {
		ch <- server.ListenAndServe()
	}()

	select {
	case err := <-ch:
		if !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("listen and Serve: %w", err)
		}
	case <-ctx.Done():
		fmt.Println("jopa: shutdown initiated")
		shutgownContext, cancel := context.WithTimeout(
			context.Background(),
			s.config.shutdownTimeout,
		)
		defer cancel()

		if err := server.Shutdown(shutgownContext); err != nil {
			_ = server.Close()

			return fmt.Errorf("shutdown http server %w", err)
		}

		fmt.Println("jopa: shutdown completed successfully")
	}

	return nil
}
