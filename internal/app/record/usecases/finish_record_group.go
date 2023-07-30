package usecases

import "github.com/victoraldir/cutcast/internal/app/record/domain"

type FinishRecordGroupUseCase interface {
	Execute(id string) error
}

type FinishRecordGroup struct {
	dbRepository        domain.RecordDbRepository
	fsWatcherRepository domain.FsWatcherRepository
}

func NewFinishRecordGroup(dbRepository domain.RecordDbRepository, fsWatcherRepository domain.FsWatcherRepository) FinishRecordGroup {
	return FinishRecordGroup{
		dbRepository:        dbRepository,
		fsWatcherRepository: fsWatcherRepository,
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

	// Unwatch
	// err = i.fsWatcherRepository.Unwatch(recordGroup.GetFullPathDirectory())

	if err != nil {
		return err
	}

	return nil
}
