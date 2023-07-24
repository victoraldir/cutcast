package usecases

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/victoraldir/cutcast/internal/app/record/domain"
)

const (
	StatusCreated    = "created"
	StatusFailed     = "failed"
	StatusProcessing = "processing"
	StatusFinished   = "finished"
	StatusCanceled   = "canceled"
	StatusError      = "error"
)

type RecordGroupCommand struct {
	Url string `json:"url"`
}

type RecordGroupResponse struct {
	Id       string `json:"id"`
	FilePath string `json:"file_path"`
	Status   string `json:"status"`
}

type CreateRecordGroupUseCase interface {
	Execute(command RecordGroupCommand) (*RecordGroupResponse, error)
}

type CreateRecordGroup struct {
	fileRepository domain.RecordFileRepository
	dbRepository   domain.RecordDbRepository
}

func NewCreateRecordGroup(fileRepository domain.RecordFileRepository, dbRepository domain.RecordDbRepository) CreateRecordGroup {
	return CreateRecordGroup{
		fileRepository: fileRepository,
		dbRepository:   dbRepository,
	}
}

func (i CreateRecordGroup) Execute(command RecordGroupCommand) (*RecordGroupResponse, error) {

	id := uuid.New().String()
	doneCh := make(chan struct{}, 1)
	videoCh := make(chan domain.Record, 1)

	mediaPath := fmt.Sprintf("/tmp/%s", id)

	recordGroup := domain.Record{
		Url:    command.Url,
		Id:     id,
		Done:   &doneCh,
		Status: domain.RecordStatusProgress,
	}

	videoCh <- recordGroup

	// Create record filesystem
	err := i.fileRepository.Create(doneCh, videoCh, mediaPath)
	if err != nil {
		return nil, err
	}

	// Create HLS filesystem
	err = i.fileRepository.CreateHLS(mediaPath, 10)
	if err != nil {
		return nil, err
	}

	recordGroup, err = i.dbRepository.Create(recordGroup)
	if err != nil {
		doneCh <- struct{}{} // Finish goroutine if error
		return nil, err
	}

	return &RecordGroupResponse{
		Id:       recordGroup.Id,
		Status:   string(recordGroup.Status),
		FilePath: mediaPath,
	}, nil
}
