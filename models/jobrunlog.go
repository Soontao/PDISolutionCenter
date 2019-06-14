package models

// JobRunLog type
type JobRunLog struct {
	BaseModel
	Running    bool
	Successful bool
	LogItems   []*Log
}
