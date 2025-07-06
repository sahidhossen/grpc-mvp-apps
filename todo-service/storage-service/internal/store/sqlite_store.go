package store

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/sahidhossen/todo/storage-service/internal/domain"
)

// ensure SQLiteStore implements the Store interface
var _ Store = (*SQLiteStore)(nil)

// SQLiteStore is an implementation of the Store interface using SQLite.
type SQLiteStore struct {
	db     *sql.DB
	logger *slog.Logger
}

// NewSQLiteStore creates a new SQLiteStore instance.
func NewSQLiteStore(db *sql.DB, logger *slog.Logger) *SQLiteStore {
	return &SQLiteStore{
		db:     db,
		logger: logger,
	}
}

// SaveTask save or update a task to the database.
func (s *SQLiteStore) SaveTask(ctx context.Context, task *domain.Task) error {
	if task.ID == "" {
		task.ID = uuid.New().String()
		task.CreatedAt = time.Now()
		task.UpdatedAt = time.Now()
		query := `INSERT INTO tasks (id, title, description, completed, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`
		_, err := s.db.ExecContext(ctx, query, task.ID, task.Title, task.Description, task.Completed, task.CreatedAt, task.UpdatedAt)
		if err != nil {
			return fmt.Errorf("failed to insert task: %w", err)
		}
		s.logger.Debug("Task inserted", "id", task.ID)
	} else {
		task.UpdatedAt = time.Now()
		query := `UPDATE tasks SET title = ?, description = ?, completed = ?, updated_at = ? WHERE id = ?`
		result, err := s.db.ExecContext(ctx, query, task.Title, task.Description, task.Completed, task.UpdatedAt, task.ID)
		if err != nil {
			return fmt.Errorf("failed to update task: %w", err)
		}
		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			return fmt.Errorf("task with ID %s not found for update", task.ID)
		}
		s.logger.Debug("Task updated", "id", task.ID)
	}
	return nil
}

// GetTask retrieves a task by its ID.
func (s *SQLiteStore) GetTask(ctx context.Context, id string) (*domain.Task, error) {
	query := `SELECT id, title, description, completed, created_at, updated_at FROM tasks WHERE id = ?`
	task := &domain.Task{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("task with ID %s not found: %w", id, err)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get task: %w", err)
	}
	return task, nil
}

// ListTasks retrieves all tasks.
func (s *SQLiteStore) ListTasks(ctx context.Context) ([]*domain.Task, error) {
	query := `SELECT * FROM tasks ORDER BY created_at DESC`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list tasks: %w", err)
	}
	defer rows.Close()

	var tasks []*domain.Task
	for rows.Next() {
		task := &domain.Task{}
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan task row: %w", err)
		}
		s.logger.Info("Task retrieved", "Completed", task.Completed)
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}
	return tasks, nil
}

// CompleteTask marks a task as completed.
func (s *SQLiteStore) ToggleTaskCompletion(ctx context.Context, id string) (*domain.Task, error) {
	query := `UPDATE tasks SET completed = NOT completed, updated_at = ? WHERE id = ?`
	result, err := s.db.ExecContext(ctx, query, time.Now(), id)
	if err != nil {
		return nil, fmt.Errorf("failed to complete task: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, fmt.Errorf("task with ID %s not found for completion", id)
	}

	return s.GetTask(ctx, id)
}

// GetTaskStats retrieves the total, completed, and remaining task counts.
func (s *SQLiteStore) GetTaskStats(ctx context.Context) (*domain.TaskStats, error) {
	query := `
		SELECT
			COUNT(*) AS total,
			COALESCE(SUM(CASE WHEN completed THEN 1 ELSE 0 END), 0) AS completed
		FROM tasks;
	`
	stats := &domain.TaskStats{}
	err := s.db.QueryRowContext(ctx, query).Scan(&stats.Total, &stats.Completed)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve task stats: %w", err)
	}

	stats.Pending = stats.Total - stats.Completed // Calculate pending tasks

	s.logger.Debug("Retrieved task stats", "total", stats.Total, "completed", stats.Completed, "pending", stats.Pending)
	return stats, nil
}
