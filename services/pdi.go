// Package services apis
// PDI Package Services
package services

import (
	"github.com/jinzhu/gorm"
)

// PDIService type
type PDIService struct {
	db       *gorm.DB
	services *Services
}
