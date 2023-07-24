package domain

import "strings"

type Trim struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	RecordId  string `json:"record_id"`
}

func (t Trim) IsValid() bool {
	return t.StartTime != "" && t.EndTime != ""
}

func (t Trim) GetStartEndTimeFormatted() string {
	return strings.Replace(t.StartTime, ":", "", -1) + "_" + strings.Replace(t.EndTime, ":", "", -1)
}
