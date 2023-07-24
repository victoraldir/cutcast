package adapters

import "github.com/victoraldir/cutcast/internal/app/record/domain"

type TrimDbMemoryRepository struct {
	Trims map[string][]domain.Trim
}

func NewTrimDbMemoryRepository() TrimDbMemoryRepository {
	return TrimDbMemoryRepository{
		Trims: make(map[string][]domain.Trim, 100),
	}
}

func (r TrimDbMemoryRepository) Create(recordId string, trim domain.Trim) (domain.Trim, error) {
	r.Trims[recordId] = append(r.Trims[recordId], trim)
	return trim, nil
}

func (r TrimDbMemoryRepository) List(recordId string) ([]domain.Trim, error) {
	return r.Trims[recordId], nil
}
