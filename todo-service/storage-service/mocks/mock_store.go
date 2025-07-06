package mocks

import (
	"context"

	"github.com/sahidhossen/todo/storage-service/internal/domain"
	"github.com/stretchr/testify/mock"
)

type MockStore struct {
	mock.Mock
}

func (m *MockStore) SaveTask(ctx context.Context, task *domain.Task) error {
	if task.ID == "" {
		task.ID = "mock-task-id-" + task.Title // More deterministic ID for tests
	}
	args := m.Called(ctx, task)
	return args.Error(0)
}
func (m *MockStore) GetTask(ctx context.Context, id string) (*domain.Task, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Task), args.Error(1)
}
func (m *MockStore) ListTasks(ctx context.Context) ([]*domain.Task, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Task), args.Error(1)
}
func (m *MockStore) ToggleTaskCompletion(ctx context.Context, id string) (*domain.Task, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Task), args.Error(1)
}
func (m *MockStore) GetTaskStats(ctx context.Context) (*domain.TaskStats, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.TaskStats), args.Error(1)
}
