package adapters

import "github.com/victoraldir/cutcast/internal/app/record/domain"

type RecordDbMemoryRepository struct {
	Records map[string]domain.Record
}

func NewRecordDbMemoryRepository() RecordDbMemoryRepository {
	return RecordDbMemoryRepository{
		Records: make(map[string]domain.Record, 100),
	}
}

func (r RecordDbMemoryRepository) Create(record domain.Record) (domain.Record, error) {
	r.Records[record.Id] = record
	return record, nil
}

func (r RecordDbMemoryRepository) Update(record domain.Record) (domain.Record, error) {
	r.Records[record.Id] = record
	return record, nil
}

func (r RecordDbMemoryRepository) Find(id string) (domain.Record, error) {
	return r.Records[id], nil
}

func (r RecordDbMemoryRepository) List() ([]domain.Record, error) {
	records := make([]domain.Record, 0)
	for _, record := range r.Records {
		records = append(records, record)
	}
	return records, nil
}
