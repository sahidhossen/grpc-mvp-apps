package store

import (
	"context"

	"github.com/sahidhossen/todo/storage-service/internal/domain"
)

// Store interface with method signature
type Store interface {
	SaveTask(ctx context.Context, task *domain.Task) error
	GetTask(ctx context.Context, id string) (*domain.Task, error)
	ListTasks(ctx context.Context) ([]*domain.Task, error)
	ToggleTaskCompletion(ctx context.Context, id string) (*domain.Task, error)
	GetTaskStats(ctx context.Context) (*domain.TaskStats, error)
}
