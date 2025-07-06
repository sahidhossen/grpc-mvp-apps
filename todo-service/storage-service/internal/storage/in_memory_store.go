/*
*
deprecated
*
*/
package storage

import (
	"log/slog"
	"sync"

	"github.com/google/uuid"
)

type Task struct {
	ID          string
	Title       string
	Description string
	Completed   bool
}

type InMemoryStore struct {
	mu     sync.RWMutex
	tasks  map[string]Task
	logger *slog.Logger
}

func NewInMemoryStore(logger *slog.Logger) *InMemoryStore {
	if logger == nil {
		logger = slog.Default()
	}
	return &InMemoryStore{
		tasks:  make(map[string]Task),
		logger: logger,
	}
}

func (s *InMemoryStore) Create(title, desc string) Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := uuid.New().String()
	t := Task{ID: id, Title: title, Description: desc, Completed: false}
	s.tasks[id] = t
	s.logger.Info("Task created", "id", id, "title", title)
	return t
}

func (s *InMemoryStore) List() []Task {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		list = append(list, t)
	}
	s.logger.Debug("Listed all tasks", "count", len(list))
	return list
}

func (s *InMemoryStore) Get(id string) (Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	t, ok := s.tasks[id]
	s.logger.Debug("Attempted to get task", "id", id, "found", ok)
	return t, ok
}

func (s *InMemoryStore) Complete(id string) (Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	t, ok := s.tasks[id]
	if !ok {
		s.logger.Warn("Attempted to complete non-existent task", "id", id)
		return Task{}, false
	}
	if t.Completed {
		s.logger.Info("Task already completed", "id", id)
		return t, true
	}
	t.Completed = true
	s.tasks[id] = t
	s.logger.Info("Task completed", "id", id)
	return t, true
}
