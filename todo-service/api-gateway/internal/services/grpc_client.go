package services

import (
	"context"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/sahidhossen/todo/proto/task_service" // Alias for generated code
)

// GRPCClient wraps the gRPC client for the TaskService.
type GRPCClient struct {
	client pb.TaskServiceClient
	conn   *grpc.ClientConn
	logger *slog.Logger
}

// TaskService interface for interacting with the Task gRPC service.
type TaskService interface {
	CreateTask(ctx context.Context, title, description string) (*pb.Task, error)
	GetTask(ctx context.Context, id string) (*pb.Task, error)
	ListTasks(ctx context.Context) ([]*pb.Task, error)
	ToggleTaskCompletion(ctx context.Context, id string) (*pb.Task, error)
	GetTaskStats(ctx context.Context) (*pb.GetTaskStatsResponse, error) // NEW: Add this
	Close() error
}

// NewGRPCClient creates a new GRPCClient and establishes a gRPC connection.
func NewGRPCClient(addr string, logger *slog.Logger) (TaskService, error) {
	if logger == nil {
		logger = slog.Default()
	}

	// In a real application, you'd use credentials.NewClientTLSFromFile or similar
	// For simplicity, we use insecure transport.
	logger.Info("Connecting to gRPC storage service", "address", addr)
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.Error("Failed to connect to gRPC server", "address", addr, "error", err)
		return nil, err
	}

	client := pb.NewTaskServiceClient(conn)
	logger.Info("Successfully connected to gRPC server!")
	return &GRPCClient{
		client: client,
		conn:   conn,
		logger: logger,
	}, nil
}

// Close method closes the gRPC connection.
func (c *GRPCClient) Close() error {
	c.logger.Info("Closing gRPC connection to storage service")
	return c.conn.Close()
}

// CreateTask calls the gRPC CreateTask method.
func (c *GRPCClient) CreateTask(ctx context.Context, title, description string) (*pb.Task, error) {
	resp, err := c.client.CreateTask(ctx, &pb.CreateTaskRequest{Title: title, Description: description})
	if err != nil {
		c.logger.Error("gRPC CreateTask failed", "error", err)
		return nil, err
	}
	return resp.Task, nil
}

// GetTask calls the gRPC GetTask method.
func (c *GRPCClient) GetTask(ctx context.Context, id string) (*pb.Task, error) {
	resp, err := c.client.GetTask(ctx, &pb.GetTaskRequest{Id: id})
	if err != nil {
		c.logger.Error("gRPC GetTask failed", "id", id, "error", err)
		return nil, err
	}
	return resp.Task, nil
}

// ListTasks calls the gRPC ListTasks method.
func (c *GRPCClient) ListTasks(ctx context.Context) ([]*pb.Task, error) {
	resp, err := c.client.ListTasks(ctx, &pb.ListTasksRequest{})
	if err != nil {
		c.logger.Error("gRPC ListTasks failed", "error", err)
		return nil, err
	}
	return resp.Tasks, nil
}

// CompleteTask calls the gRPC CompleteTask method.
func (c *GRPCClient) ToggleTaskCompletion(ctx context.Context, id string) (*pb.Task, error) {
	resp, err := c.client.ToggleTaskCompletion(ctx, &pb.ToggleTaskCompletionRequest{Id: id})
	if err != nil {
		c.logger.Error("gRPC CompleteTask failed", "id", id, "error", err)
		return nil, err
	}
	return resp.Task, nil
}

// GetTaskStats calls the gRPC GetTaskStats method.
func (c *GRPCClient) GetTaskStats(ctx context.Context) (*pb.GetTaskStatsResponse, error) {
	resp, err := c.client.GetTaskStats(ctx, &pb.GetTaskStatsRequest{})
	if err != nil {
		c.logger.Error("gRPC GetTaskStats failed", "error", err)
		return nil, err
	}
	return resp, nil
}
