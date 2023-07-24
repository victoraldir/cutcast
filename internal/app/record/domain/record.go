package domain

type Status string

const (
	RecordStatusProgress = "progress"
	RecordStatusDone     = "done"
	RecordStatusError    = "error"
)

type Record struct {
	Id     string
	Status Status
	Url    string
	Done   *chan struct{}
}

func (r Record) GetFullPathDirectory() string {

	if r.Id == "" {
		panic("Record Id is required")
	}

	return "/tmp/" + r.Id
}

func (s Status) String() string {
	return string(s)
}
