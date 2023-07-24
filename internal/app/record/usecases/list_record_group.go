package usecases

import "github.com/victoraldir/cutcast/internal/app/record/domain"

type ListRecordGroupUseCase interface {
	Execute() ([]RecordGroupResponse, error)
}

type ListRecordGroup struct {
	recordGroupRepository domain.RecordDbRepository
}

func NewListRecordGroup(recordGroupRepository domain.RecordDbRepository) ListRecordGroup {
	return ListRecordGroup{
		recordGroupRepository: recordGroupRepository,
	}
}

func (i ListRecordGroup) Execute() ([]RecordGroupResponse, error) {
	recordGroups, err := i.recordGroupRepository.List()
	if err != nil {
		return nil, err
	}

	recordGroupResponses := make([]RecordGroupResponse, 0)
	for _, recordGroup := range recordGroups {
		recordGroupResponses = append(recordGroupResponses, RecordGroupResponse{
			Id:       recordGroup.Id,
			FilePath: recordGroup.GetFullPathDirectory(),
			Status:   recordGroup.Status.String(),
		})
	}

	return recordGroupResponses, nil
}
