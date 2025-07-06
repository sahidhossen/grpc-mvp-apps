package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sahidhossen/todo/api-gateway/internal/middleware"
)

// HTTPService represents the HTTP server for the API Gateway.
type HTTPService struct {
	srv    *http.Server
	logger *slog.Logger
}

// NewHTTPService creates a new HTTPService.
func NewHTTPService(port string, router *mux.Router, logger *slog.Logger) *HTTPService {
	if logger == nil {
		logger = slog.Default()
	}
	return &HTTPService{
		srv: &http.Server{
			Addr:         ":" + port,
			Handler:      router,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
		logger: logger,
	}
}

// Start begins listening for HTTP requests.
func (s *HTTPService) Start() error {
	s.logger.Info("REST server listening", "port", s.srv.Addr)
	return s.srv.ListenAndServe()
}

// Shutdown gracefully shuts down the HTTP server.
func (s *HTTPService) Shutdown(ctx context.Context) error {
	s.logger.Info("Shutting down REST server...")
	return s.srv.Shutdown(ctx)
}

// NewRouter initializes and returns a new Gorilla Mux router with common middleware.
func NewRouter(logger *slog.Logger) *mux.Router {
	if logger == nil {
		logger = slog.Default()
	}
	router := mux.NewRouter()

	// Add global middleware (e.g., logging, CORS)
	router.Use(middleware.LoggingMiddleware(logger))
	router.Use(middleware.CORSMiddleware(logger))

	return router
}
