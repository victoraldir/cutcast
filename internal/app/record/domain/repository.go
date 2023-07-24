package domain

type RecordFileRepository interface {
	Create(done <-chan struct{}, record <-chan Record, mediaDir string) error
	Trim(id string, trim Trim, mediaDir string) (*string, error)
	CreateHLS(mediaDir string, segmentDuration int) error
}

type RecordDbRepository interface {
	Create(record Record) (Record, error)
	Update(record Record) (Record, error)
	Find(id string) (Record, error)
	List() ([]Record, error)
}

type TrimDbRepository interface {
	Create(recordId string, trim Trim) (Trim, error)
	List(recordId string) ([]Trim, error)
}
