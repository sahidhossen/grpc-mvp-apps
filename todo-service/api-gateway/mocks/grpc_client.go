package mocks

import (
	"context"

	pb "github.com/sahidhossen/todo/proto/task_service"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

// MockTaskServiceClient is a mock implementation of pb.TaskServiceClient
type MockTaskServiceClient struct {
	mock.Mock
}

func (m *MockTaskServiceClient) CreateTask(ctx context.Context, in *pb.CreateTaskRequest, opts ...grpc.CallOption) (*pb.CreateTaskResponse, error) {
	args := m.Called(ctx, in)
	// Check if the first argument (response) is nil before type asserting
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*pb.CreateTaskResponse), args.Error(1)
}

func (m *MockTaskServiceClient) GetTask(ctx context.Context, in *pb.GetTaskRequest, opts ...grpc.CallOption) (*pb.GetTaskResponse, error) {
	args := m.Called(ctx, in)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*pb.GetTaskResponse), args.Error(1)
}

func (m *MockTaskServiceClient) ListTasks(ctx context.Context, in *pb.ListTasksRequest, opts ...grpc.CallOption) (*pb.ListTasksResponse, error) {
	args := m.Called(ctx, in)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*pb.ListTasksResponse), args.Error(1)
}

func (m *MockTaskServiceClient) ToggleTaskCompletion(ctx context.Context, in *pb.ToggleTaskCompletionRequest, opts ...grpc.CallOption) (*pb.ToggleTaskCompletionResponse, error) {
	args := m.Called(ctx, in)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*pb.ToggleTaskCompletionResponse), args.Error(1)
}

func (m *MockTaskServiceClient) CompleteTask(ctx context.Context, in *pb.CompleteTaskRequest, opts ...grpc.CallOption) (*pb.CompleteTaskResponse, error) {
	args := m.Called(ctx, in)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*pb.CompleteTaskResponse), args.Error(1)
}

func (m *MockTaskServiceClient) GetTaskStats(ctx context.Context, in *pb.GetTaskStatsRequest, opts ...grpc.CallOption) (*pb.GetTaskStatsResponse, error) {
	args := m.Called(ctx, in)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*pb.GetTaskStatsResponse), args.Error(1)
}
