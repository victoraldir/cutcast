package usecases

import (
	"fmt"

	"github.com/victoraldir/cutcast/internal/app/record/domain"
)

type TrimRecordGroupCommand struct {
	RecordId  string `json:"record_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type TrimRecordGroupResponse struct {
	RecordId  string `json:"record_id"`
	FilePath  string `json:"file_path"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type TrimRecordGroupUseCase interface {
	Execute(command TrimRecordGroupCommand) (*TrimRecordGroupResponse, error)
}

type TrimRecordGroup struct {
	fileRepository            domain.RecordFileRepository
	trimRecordGroupRepository domain.TrimDbRepository
}

func NewTrimRecordGroup(fileRepository domain.RecordFileRepository,
	trimRecordGroupRepository domain.TrimDbRepository) TrimRecordGroup {
	return TrimRecordGroup{
		fileRepository:            fileRepository,
		trimRecordGroupRepository: trimRecordGroupRepository,
	}
}

func (i TrimRecordGroup) Execute(command TrimRecordGroupCommand) (*TrimRecordGroupResponse, error) {

	trim := domain.Trim{
		StartTime: command.StartTime,
		EndTime:   command.EndTime,
	}

	mediaDir := fmt.Sprintf("/tmp/%s", command.RecordId)

	filePath, err := i.fileRepository.Trim(command.RecordId, trim, mediaDir)
	if err != nil {
		return nil, err
	}

	_, err = i.trimRecordGroupRepository.Create(command.RecordId, trim)
	if err != nil {
		return nil, err
	}

	return &TrimRecordGroupResponse{
		RecordId: command.RecordId,
		FilePath: *filePath,
	}, nil
}
