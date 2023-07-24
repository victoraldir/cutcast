package usecases

import "github.com/victoraldir/cutcast/internal/app/record/domain"

type ListTrimRecordGroupUseCase interface {
	Execute(recordId string) ([]TrimRecordGroupResponse, error)
}

type ListTrimRecordGroup struct {
	trimRecordGroupRepository domain.TrimDbRepository
}

func NewListTrimRecordGroup(trimRecordGroupRepository domain.TrimDbRepository) *ListTrimRecordGroup {
	return &ListTrimRecordGroup{
		trimRecordGroupRepository: trimRecordGroupRepository,
	}
}

func (i *ListTrimRecordGroup) Execute(id string) ([]TrimRecordGroupResponse, error) {
	trimRecordGroups, err := i.trimRecordGroupRepository.List(id)
	if err != nil {
		return nil, err
	}

	trimRecordGroupResponses := make([]TrimRecordGroupResponse, 0)
	for _, trimRecordGroup := range trimRecordGroups {
		trimRecordGroupResponses = append(trimRecordGroupResponses, TrimRecordGroupResponse{
			StartTime: trimRecordGroup.StartTime,
			EndTime:   trimRecordGroup.EndTime,
			RecordId:  trimRecordGroup.RecordId,
		})
	}

	return trimRecordGroupResponses, nil
}
