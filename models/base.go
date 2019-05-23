package models

import (
	"github.com/jinzhu/gorm"
)

// BaseModel type
type BaseModel struct {
	gorm.Model
	Name        string
	Description string `gorm:"size:2048"`
	CreatedBy   *User
	UpdatedBy   *User
}
