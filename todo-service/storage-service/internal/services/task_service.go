package services

import (
	"context"
	"log/slog"

	"github.com/sahidhossen/todo/storage-service/internal/converters"
	"github.com/sahidhossen/todo/storage-service/internal/domain"
	"github.com/sahidhossen/todo/storage-service/internal/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/sahidhossen/todo/proto/task_service"
)

// TaskServiceServer implements the gRPC server interface for TaskService.
type TaskServiceServer struct {
	pb.UnimplementedTaskServiceServer // Must be embedded for forward compatibility
	store                             store.Store
	logger                            *slog.Logger
}

// NewTaskServiceServer creates a new TaskServiceServer.
func NewTaskServiceServer(store store.Store, logger *slog.Logger) *TaskServiceServer {
	if logger == nil {
		logger = slog.Default()
	}
	return &TaskServiceServer{
		store:  store,
		logger: logger,
	}
}

// CreateTask handles the gRPC request to create a new task.
func (s *TaskServiceServer) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	if req.Title == "" {
		s.logger.Warn("CreateTask request missing title")
		return nil, status.Errorf(codes.InvalidArgument, "title cannot be empty")
	}

	domainTask := &domain.Task{
		Title:       req.Title,
		Description: req.Description,
		Completed:   false,
	}

	err := s.store.SaveTask(ctx, domainTask)
	if err != nil {
		s.logger.Error("Failed to save task to store", "error", err)
		return nil, status.Errorf(codes.Internal, "failed to save task: %v", err)
	}

	pbTaskResponse := converters.DomainToProtoTask(domainTask)
	return &pb.CreateTaskResponse{Task: pbTaskResponse}, nil
}

// GetTask handles the gRPC request to get a task by ID.
func (s *TaskServiceServer) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.GetTaskResponse, error) {
	if req.Id == "" {
		s.logger.Warn("GetTask request missing ID")
		return nil, status.Errorf(codes.InvalidArgument, "task ID cannot be empty")
	}

	task, err := s.store.GetTask(ctx, req.Id)
	if err != nil {
		s.logger.Warn("gRPC: Task not found", "id", req.Id)
		return nil, status.Errorf(codes.NotFound, "task with ID %s not found", req.Id)
	}
	s.logger.Info("gRPC: Task retrieved", "id", task.ID)
	return &pb.GetTaskResponse{
		Task: converters.DomainToProtoTask(task),
	}, nil
}

// ListTasks handles the gRPC request to list all tasks.
func (s *TaskServiceServer) ListTasks(ctx context.Context, req *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	tasks, err := s.store.ListTasks(ctx)
	if err != nil {
		s.logger.Error("Failed to list tasks from store", "error", err)
		return nil, status.Errorf(codes.Internal, "failed to list tasks: %v", err)
	}
	pbTasks := make([]*pb.Task, len(tasks))
	for i, task := range tasks {
		pbTasks[i] = converters.DomainToProtoTask(task)
	}
	s.logger.Info("gRPC: Listed tasks", "count", len(pbTasks))
	return &pb.ListTasksResponse{Tasks: pbTasks}, nil
}

// ToggleTaskCompletion handles the gRPC request to mark a task as completed.
func (s *TaskServiceServer) ToggleTaskCompletion(ctx context.Context, req *pb.ToggleTaskCompletionRequest) (*pb.ToggleTaskCompletionResponse, error) {

	s.logger.Info("gRPC: Received ToggleTaskCompletion request", "id", req.Id)

	if req.Id == "" {
		s.logger.Warn("Task ID missing!")
		return nil, status.Errorf(codes.InvalidArgument, "task ID cannot be empty")
	}

	task, err := s.store.ToggleTaskCompletion(ctx, req.Id)
	if err != nil {
		s.logger.Warn("gRPC: Task not found for changes", "id", req.Id)
		return nil, status.Errorf(codes.NotFound, "task with ID %s not found", req.Id)
	}

	return &pb.ToggleTaskCompletionResponse{
		Task: converters.DomainToProtoTask(task),
	}, nil
}

// GetTaskStats implements the gRPC GetTaskStats method.
func (s *TaskServiceServer) GetTaskStats(ctx context.Context, req *pb.GetTaskStatsRequest) (*pb.GetTaskStatsResponse, error) {
	s.logger.Info("Received GetTaskStats request")

	stats, err := s.store.GetTaskStats(ctx)
	if err != nil {
		s.logger.Error("Failed to retrieve task stats from store", "error", err)
		return nil, status.Errorf(codes.Internal, "failed to retrieve task stats: %v", err)
	}

	return &pb.GetTaskStatsResponse{
		TotalTasks:     stats.Total,
		CompletedTasks: stats.Completed,
		PendingTasks:   stats.Pending,
	}, nil
}
