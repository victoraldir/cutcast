package usecases

import "github.com/victoraldir/cutcast/internal/app/record/domain"

type CreateHLSRecordGroupCommand struct {
	Path string
}

type CreateHLSRecordGroupResponse struct {
	WatcherId string `json:"watcher_id"`
}

type CreateHLSRecordGroupUseCase interface {
	Execute(command CreateHLSRecordGroupCommand) (*CreateHLSRecordGroupResponse, error)
}

type createHLSRecordGroupUseCase struct {
	fileRepository domain.RecordFileRepository
}

func NewCreateHLSRecordGroupUseCase(fileRepository domain.RecordFileRepository) CreateHLSRecordGroupUseCase {
	return &createHLSRecordGroupUseCase{
		fileRepository: fileRepository,
	}
}

func (i *createHLSRecordGroupUseCase) Execute(command CreateHLSRecordGroupCommand) (*CreateHLSRecordGroupResponse, error) {

	err := i.fileRepository.CreateHLS(command.Path, 2)
	if err != nil {
		return nil, err
	}

	return &CreateHLSRecordGroupResponse{}, nil
}
