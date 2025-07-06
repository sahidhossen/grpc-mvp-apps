package services

import (
	"context"
	"errors"
	"log/slog"
	"testing"

	"github.com/sahidhossen/todo/api-gateway/mocks"
	pb "github.com/sahidhossen/todo/proto/task_service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGRPCClient_CreateTask_Success(t *testing.T) {
	mockClient := new(mocks.MockTaskServiceClient)
	grpcClient := &GRPCClient{
		client: mockClient,
		logger: slog.Default(),
	}

	expectedTask := &pb.Task{Id: "123", Title: "Test Task", Description: "Description", Completed: false}
	mockClient.On("CreateTask", mock.Anything, &pb.CreateTaskRequest{Title: "Test Task", Description: "Description"}).
		Return(&pb.CreateTaskResponse{Task: expectedTask}, nil)

	task, err := grpcClient.CreateTask(context.Background(), "Test Task", "Description")

	assert.NoError(t, err)
	assert.Equal(t, expectedTask, task)
	mockClient.AssertExpectations(t) // Verify that mock method was called
}

func TestGRPCClient_CreateTask_Error(t *testing.T) {
	mockClient := new(mocks.MockTaskServiceClient)
	grpcClient := &GRPCClient{
		client: mockClient,
		logger: slog.Default(),
	}

	expectedErr := status.Error(codes.Internal, "something went wrong")
	mockClient.On("CreateTask", mock.Anything, &pb.CreateTaskRequest{Title: "Error Task", Description: ""}).
		Return(nil, expectedErr) // Return nil response, and an error

	task, err := grpcClient.CreateTask(context.Background(), "Error Task", "")

	assert.Error(t, err)
	assert.Nil(t, task)
	assert.Equal(t, expectedErr, err)
	mockClient.AssertExpectations(t)
}

func TestGRPCClient_GetTask_Success(t *testing.T) {
	mockClient := new(mocks.MockTaskServiceClient)
	grpcClient := &GRPCClient{
		client: mockClient,
		logger: slog.Default(),
	}

	expectedTask := &pb.Task{Id: "456", Title: "Fetched Task", Completed: true}
	mockClient.On("GetTask", mock.Anything, &pb.GetTaskRequest{Id: "456"}).
		Return(&pb.GetTaskResponse{Task: expectedTask}, nil)

	task, err := grpcClient.GetTask(context.Background(), "456")

	assert.NoError(t, err)
	assert.Equal(t, expectedTask, task)
	mockClient.AssertExpectations(t)
}

func TestGRPCClient_GetTask_NotFound(t *testing.T) {
	mockClient := new(mocks.MockTaskServiceClient)
	grpcClient := &GRPCClient{
		client: mockClient,
		logger: slog.Default(),
	}

	expectedErr := status.Error(codes.NotFound, "task not found")
	mockClient.On("GetTask", mock.Anything, &pb.GetTaskRequest{Id: "nonexistent"}).
		Return(nil, expectedErr)

	task, err := grpcClient.GetTask(context.Background(), "nonexistent")

	assert.Error(t, err)
	assert.Nil(t, task)
	assert.Equal(t, expectedErr, err)
	mockClient.AssertExpectations(t)
}

func TestGRPCClient_ListTasks_Success(t *testing.T) {
	mockClient := new(mocks.MockTaskServiceClient)
	grpcClient := &GRPCClient{
		client: mockClient,
		logger: slog.Default(),
	}

	expectedTasks := []*pb.Task{
		{Id: "1", Title: "Task 1"},
		{Id: "2", Title: "Task 2"},
	}
	mockClient.On("ListTasks", mock.Anything, &pb.ListTasksRequest{}).
		Return(&pb.ListTasksResponse{Tasks: expectedTasks}, nil)

	tasks, err := grpcClient.ListTasks(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, expectedTasks, tasks)
	mockClient.AssertExpectations(t)
}

func TestGRPCClient_ListTasks_Error(t *testing.T) {
	mockClient := new(mocks.MockTaskServiceClient)
	grpcClient := &GRPCClient{
		client: mockClient,
		logger: slog.Default(),
	}

	expectedErr := errors.New("database connection failed") // Non-gRPC specific error
	mockClient.On("ListTasks", mock.Anything, &pb.ListTasksRequest{}).
		Return(nil, expectedErr)

	tasks, err := grpcClient.ListTasks(context.Background())

	assert.Error(t, err)
	assert.Nil(t, tasks)
	assert.Equal(t, expectedErr, err)
	mockClient.AssertExpectations(t)
}

func TestGRPCClient_ToggleTaskCompletion_Success(t *testing.T) {
	mockClient := new(mocks.MockTaskServiceClient)
	grpcClient := &GRPCClient{
		client: mockClient,
		logger: slog.Default(),
	}

	expectedTask := &pb.Task{Id: "789", Title: "Toggled Task", Completed: true}
	mockClient.On("ToggleTaskCompletion", mock.Anything, &pb.ToggleTaskCompletionRequest{Id: "789"}).
		Return(&pb.ToggleTaskCompletionResponse{Task: expectedTask}, nil)

	task, err := grpcClient.ToggleTaskCompletion(context.Background(), "789")

	assert.NoError(t, err)
	assert.Equal(t, expectedTask, task)
	mockClient.AssertExpectations(t)
}

func TestGRPCClient_ToggleTaskCompletion_Error(t *testing.T) {
	mockClient := new(mocks.MockTaskServiceClient)
	grpcClient := &GRPCClient{
		client: mockClient,
		logger: slog.Default(),
	}

	expectedErr := status.Error(codes.InvalidArgument, "invalid ID format")
	mockClient.On("ToggleTaskCompletion", mock.Anything, &pb.ToggleTaskCompletionRequest{Id: "invalid-id"}).
		Return(nil, expectedErr)

	task, err := grpcClient.ToggleTaskCompletion(context.Background(), "invalid-id")

	assert.Error(t, err)
	assert.Nil(t, task)
	assert.Equal(t, expectedErr, err)
	mockClient.AssertExpectations(t)
}
