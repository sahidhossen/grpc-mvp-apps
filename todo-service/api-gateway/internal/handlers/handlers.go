package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sahidhossen/todo/api-gateway/internal/httputil"
	"github.com/sahidhossen/todo/api-gateway/internal/services"
)

// Handler manages HTTP requests by interacting with the gRPC task service.
type Handler struct {
	taskClient services.TaskService
	logger     *slog.Logger
}

// New creates a new Handler.
func New(taskClient services.TaskService, logger *slog.Logger) *Handler {
	if logger == nil {
		logger = slog.Default()
	}
	return &Handler{
		taskClient: taskClient,
		logger:     logger,
	}
}

// RegisterRoutes sets up all the API routes for the application.
func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", h.CreateTask).Methods("POST")
	r.HandleFunc("/tasks", h.ListTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", h.GetTask).Methods("GET")
	r.HandleFunc("/tasks/{id}/toggle-task-complete", h.ToggleTaskCompletion).Methods("PATCH")
	r.HandleFunc("/stats", h.GetTaskStats).Methods("GET")
}

// CreateTask handles the creation of a new task.
func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httputil.HandleError(w, r, h.logger, err, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Title == "" {
		httputil.HandleError(w, r, h.logger, nil, "Title cannot be empty", http.StatusBadRequest)
		return
	}

	ctx, cancel := httputil.WithTimeout(r)
	defer cancel()

	task, err := h.taskClient.CreateTask(ctx, req.Title, req.Description)
	if err != nil {
		httputil.HandleGrpcError(w, r, h.logger, err, "Failed to create task")
		return
	}

	httputil.HandleSuccess(w, r, h.logger, task, http.StatusCreated)
	h.logger.Info("Task created via API", "id", task.Id, "title", task.Title)
}

// ListTasks handles listing all tasks.
func (h *Handler) ListTasks(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := httputil.WithTimeout(r)
	defer cancel()

	tasks, err := h.taskClient.ListTasks(ctx)
	if err != nil {
		httputil.HandleError(w, r, h.logger, nil, "Failed to retrieve tasks", http.StatusInternalServerError)
		return
	}

	httputil.HandleSuccess(w, r, h.logger, tasks, http.StatusOK)
	h.logger.Info("Listed tasks via API", "count", len(tasks))
}

// GetTask handles retrieving a single task by ID.
func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	ctx, cancel := httputil.WithTimeout(r)
	defer cancel()

	task, err := h.taskClient.GetTask(ctx, id)
	if err != nil {
		httputil.HandleGrpcError(w, r, h.logger, err, "Failed to retrieve task")
		return
	}

	httputil.HandleSuccess(w, r, h.logger, task, http.StatusOK)
	h.logger.Info("Task retrieved via API", "id", task.Id)
}

// ToggleTaskCompletion handles marking a task as completed.
func (h *Handler) ToggleTaskCompletion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	ctx, cancel := httputil.WithTimeout(r)
	defer cancel()

	task, err := h.taskClient.ToggleTaskCompletion(ctx, id)
	if err != nil {
		httputil.HandleGrpcError(w, r, h.logger, err, "Failed to retrieve task")
		return
	}

	httputil.HandleSuccess(w, r, h.logger, task, http.StatusOK)
	h.logger.Info("Task completed via API", "id", task.Id)
}

// GetTaskStats handles retrieving task statistics.
func (h *Handler) GetTaskStats(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := httputil.WithTimeout(r)
	defer cancel()

	stats, err := h.taskClient.GetTaskStats(ctx) // Call the gRPC client method
	if err != nil {
		httputil.HandleGrpcError(w, r, h.logger, err, "Failed to retrieve task statistics")
		return
	}

	httputil.HandleSuccess(w, r, h.logger, stats, http.StatusOK)
	h.logger.Info("Task stats retrieved via API", "total", stats.TotalTasks, "completed", stats.CompletedTasks)
}
