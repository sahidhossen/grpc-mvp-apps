package services

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"testing"

	pb "github.com/sahidhossen/todo/proto/task_service"
	"github.com/sahidhossen/todo/storage-service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewNopLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
}

func TestCreateTask_Success(t *testing.T) {
	mockStore := new(mocks.MockStore)
	service := NewTaskServiceServer(mockStore, NewNopLogger())

	req := &pb.CreateTaskRequest{
		Title:       "Test Task",
		Description: "This is a test description",
	}

	mockStore.On("SaveTask", mock.Anything, mock.AnythingOfType("*domain.Task")).Return(nil).Once()

	resp, err := service.CreateTask(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, resp.Task)
	assert.Equal(t, req.Title, resp.Task.Title)
	assert.NotEmpty(t, resp.Task.Id)
	assert.Equal(t, req.Description, resp.Task.Description)
	assert.False(t, resp.Task.Completed)

	mockStore.AssertExpectations(t)
}

func TestCreateTask_StoreError(t *testing.T) {
	mockStore := new(mocks.MockStore)
	service := NewTaskServiceServer(mockStore, NewNopLogger())

	req := &pb.CreateTaskRequest{Title: "Error Task"}

	mockStore.On("SaveTask", mock.Anything, mock.AnythingOfType("*domain.Task")).Return(errors.New("db write error")).Once()

	resp, err := service.CreateTask(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	s, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, codes.Internal, s.Code())
	assert.Equal(t, "failed to save task: db write error", s.Message())
	mockStore.AssertExpectations(t)
}
