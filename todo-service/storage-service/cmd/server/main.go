package main

import (
	"context"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "github.com/sahidhossen/todo/proto/task_service"
	"github.com/sahidhossen/todo/storage-service/internal/config"
	"github.com/sahidhossen/todo/storage-service/internal/db"
	"github.com/sahidhossen/todo/storage-service/internal/services"
	"github.com/sahidhossen/todo/storage-service/internal/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Setup structured logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg := config.LoadConfig()

	// Initialize Database Connection
	// Create directory for database file if it doesn't exist
	dbDir := "./data"
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		logger.Error("Failed to create database directory", "path", dbDir, "error", err)
		os.Exit(1)
	}

	database, err := db.NewConnection(cfg.DBPath, logger)
	if err != nil {
		logger.Error("Failed to establish database connection", "error", err)
		os.Exit(1)
	}
	defer func() {
		// Ensure the DB connection is closed
		if cerr := database.Close(); cerr != nil {
			logger.Error("Failed to close database connection", "error", cerr)
		}
	}()

	// Apply Database Schema. We can think about migration command
	if err := db.ApplySchema(database, logger); err != nil {
		logger.Error("Failed to apply database schema", "error", err)
		os.Exit(1)
	}

	taskStore := store.NewSQLiteStore(database, logger)

	// Setup gRPC server
	list, err := net.Listen("tcp", ":"+cfg.GRPCPort)

	if err != nil {
		logger.Error("Failed to listen for gRPC", "error", err)
		os.Exit(1)
	}

	server := grpc.NewServer()

	// Register task service server from gRPC
	pb.RegisterTaskServiceServer(server, services.NewTaskServiceServer(taskStore, logger))
	reflection.Register(server) // Enable gRPC reflection for debugging

	// Graceful shutdown channel
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start gRPC server in a goroutine
	go func() {
		logger.Info("gRPC server listening", "port", cfg.GRPCPort)
		if err := server.Serve(list); err != nil {
			logger.Error("gRPC server failed to serve", "error", err)
		}
	}()

	// Wait for OS signal for gracful shutdown
	sig := <-quit
	logger.Info("Shutting down gRPC server...", "signal", sig)

	// Graceful shutdown with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	stopped := make(chan struct{})
	go func() {
		server.GracefulStop() // Blocks until all active RPCs finish or context timeout
		close(stopped)
	}()

	select {

	case <-stopped:
		logger.Info("gRPC server gracefully stopped.")
	case <-ctx.Done():
		logger.Error("gRPC server did not stop gracefully in time, forcing shutdown", "error", ctx.Err())
		server.Stop()
	}

	logger.Info("Storage service exited.")

}
