package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sahidhossen/todo/api-gateway/internal/config"
	"github.com/sahidhossen/todo/api-gateway/internal/handlers"
	"github.com/sahidhossen/todo/api-gateway/internal/server"
	"github.com/sahidhossen/todo/api-gateway/internal/services"
)

func main() {
	// Setup structured logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	cfg := config.LoadConfig()

	// Initiate gRPC client for the storage service
	taskClient, err := services.NewGRPCClient(cfg.GRPCHost, logger)
	if err != nil {
		logger.Error("Faield to connect gRPC service", "error", err)
		os.Exit(1)
	}

	defer func() {
		if err := taskClient.Close(); err != nil {
			logger.Error("Failed to close gRPC service", "error", err)
		}
	}()

	//Initialize HTTP handlers with the gRPC client
	handler := handlers.New(taskClient, logger)

	//Setup Gorilla Mux router
	router := server.NewRouter(logger)
	handler.RegisterRoutes(router)

	// Register Global Fallback for OPTIONS and 404s
	router.PathPrefix("/").HandlerFunc(handlers.NotFoundHandler)

	//Create and start HTTP server
	httpService := server.NewHTTPService(cfg.Port, router, logger)

	// Graceful shutdown channel
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start HTTP server in a goroutine
	go func() {
		if err := httpService.Start(); err != nil && err != http.ErrServerClosed {
			logger.Error("HTTP server failed to listen and serve", "error", err)
			os.Exit(1)
		}
	}()

	// Wait for OS signal for graceful shutdown
	sig := <-quit
	logger.Info("Shutting down API Gateway...", "signal", sig)

	// Graceful shutdown with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := httpService.Shutdown(ctx); err != nil {
		logger.Error("API Gateway forced to shutdown", "error", err)
		os.Exit(1)
	}

	logger.Info("API Gateway gracefully stopped.")

}
