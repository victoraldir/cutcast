package usecases

import "github.com/victoraldir/cutcast/internal/app/record/domain"

type FinishRecordGroupUseCase interface {
	Execute(id string) error
}

type FinishRecordGroup struct {
	dbRepository domain.RecordDbRepository
}

func NewFinishRecordGroup(dbRepository domain.RecordDbRepository) FinishRecordGroup {
	return FinishRecordGroup{
		dbRepository: dbRepository,
	}
}

func (i FinishRecordGroup) Execute(id string) error {

	recordGroup, err := i.dbRepository.Find(id)
	if err != nil {
		return err
	}

	*recordGroup.Done <- struct{}{} // Finish goroutine

	recordGroup.Status = domain.RecordStatusDone

	_, err = i.dbRepository.Update(recordGroup)
	if err != nil {
		return err
	}

	return nil
}
