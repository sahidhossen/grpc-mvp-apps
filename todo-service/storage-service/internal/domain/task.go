package domain

import "time"

// Task represents a task in the application's core domain.
// This struct is independent of database or gRPC specific details.
type Task struct {
	ID          string
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TaskStats struct {
	Total     int32
	Completed int32
	Pending   int32
}

func (t *Task) MarkComplete() {
	if !t.Completed {
		t.Completed = true
		t.UpdatedAt = time.Now()
	}
}
