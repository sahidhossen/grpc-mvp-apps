package converters

import (
	pb "github.com/sahidhossen/todo/proto/task_service"
	"github.com/sahidhossen/todo/storage-service/internal/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// DomainToProtoTask converts a domain.Task to a pb.Task.
func DomainToProtoTask(dTask *domain.Task) *pb.Task {
	if dTask == nil {
		return nil
	}
	return &pb.Task{
		Id:          dTask.ID,
		Title:       dTask.Title,
		Description: dTask.Description,
		Completed:   dTask.Completed,
		CreatedAt:   timestamppb.New(dTask.CreatedAt),
		UpdatedAt:   timestamppb.New(dTask.UpdatedAt),
	}
}

// ProtoToDomainTask converts a pb.Task to a domain.Task.
func ProtoToDomainTask(pTask *pb.Task) *domain.Task {
	if pTask == nil {
		return nil
	}
	return &domain.Task{
		ID:          pTask.GetId(),
		Title:       pTask.GetTitle(),
		Description: pTask.GetDescription(),
		Completed:   pTask.GetCompleted(),
		CreatedAt:   pTask.GetCreatedAt().AsTime(),
		UpdatedAt:   pTask.GetUpdatedAt().AsTime(),
	}
}
