package models

// JobRunLog type
type JobRunLog struct {
	BaseModel
	Running    bool
	Successful bool
	Result     string `gorm:"size:10240"`
}
