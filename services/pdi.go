// Package services apis
// PDI Package Services
package services

import (
	"github.com/jinzhu/gorm"
	"github.com/magic003/alice"
)

// PDIService type
type PDIService struct {
	alice.BaseModule
	DB *gorm.DB
}

func NewPDIService(db *gorm.DB) *PDIService {
	return &PDIService{DB: db}
}
