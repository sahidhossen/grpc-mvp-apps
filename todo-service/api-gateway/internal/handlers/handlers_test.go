package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/sahidhossen/todo/api-gateway/mocks"
	pb "github.com/sahidhossen/todo/proto/task_service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func mockWithTimeout(r *http.Request) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

var WithTimeout = mockWithTimeout

func newTestRequest(method, path string, body interface{}) *http.Request {
	var reqBody io.Reader
	if body != nil {
		jsonBytes, _ := json.Marshal(body)
		reqBody = bytes.NewBuffer(jsonBytes)
	}
	req := httptest.NewRequest(method, path, reqBody)
	if reqBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req
}

// decodeResponse decodes a JSON response body into a target struct.
func decodeResponse(rr *httptest.ResponseRecorder, target interface{}) error {
	return json.NewDecoder(rr.Body).Decode(target)
}

func TestCreateTask_Success(t *testing.T) {
	mockTaskClient := new(mocks.MockTaskService)
	handler := &Handler{
		taskClient: mockTaskClient,
		logger:     slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
	}

	// Mock the gRPC client's CreateTask call
	reqBody := struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}{
		Title:       "New Task",
		Description: "Task description",
	}
	expectedTask := &pb.Task{Id: "task1", Title: "New Task", Description: "Task description", Completed: false}

	mockTaskClient.On("CreateTask", mock.AnythingOfType("*context.timerCtx"), reqBody.Title, reqBody.Description).
		Return(expectedTask, nil).Once()

	req := newTestRequest(http.MethodPost, "/tasks", reqBody)
	rr := httptest.NewRecorder()

	handler.CreateTask(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var actualTask pb.Task
	err := decodeResponse(rr, &actualTask)
	assert.NoError(t, err)
	assert.Equal(t, expectedTask.Id, actualTask.Id)
	assert.Equal(t, expectedTask.Title, actualTask.Title)
	assert.Equal(t, expectedTask.Description, actualTask.Description)
	assert.Equal(t, expectedTask.Completed, actualTask.Completed)

	mockTaskClient.AssertExpectations(t) // Verify all mocked methods were called
}
