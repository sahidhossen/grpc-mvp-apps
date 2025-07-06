package mocks

import (
	"context"

	pb "github.com/sahidhossen/todo/proto/task_service"
	"github.com/stretchr/testify/mock"
)

type MockTaskService struct {
	mock.Mock
}

// Ensure MockTaskService implements services.TaskService
// var _ services.TaskService = &MockTaskService{}

func (m *MockTaskService) CreateTask(ctx context.Context, title, description string) (*pb.Task, error) {
	args := m.Called(ctx, title, description)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*pb.Task), args.Error(1)
}

func (m *MockTaskService) GetTask(ctx context.Context, id string) (*pb.Task, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*pb.Task), args.Error(1)
}

func (m *MockTaskService) ListTasks(ctx context.Context) ([]*pb.Task, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*pb.Task), args.Error(1)
}

func (m *MockTaskService) ToggleTaskCompletion(ctx context.Context, id string) (*pb.Task, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*pb.Task), args.Error(1)
}

func (m *MockTaskService) GetTaskStats(ctx context.Context) (*pb.GetTaskStatsResponse, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*pb.GetTaskStatsResponse), args.Error(1)
}

func (m *MockTaskService) Close() error {
	args := m.Called()
	return args.Error(0)
}
