package services

import (
	"github.com/jinzhu/gorm"
	"github.com/magic003/alice"
)

// Service interface
type Service interface {
	init()
}

// Services type
type Services struct {
	alice.BaseModule
	PDIService      *PDIService
	ScheduleService *ScheduleService
}

// NewServices constructor
func NewServices(db *gorm.DB) (*Services, error) {

	services := &Services{}

	pdiService := &PDIService{}

	scheduleService := &ScheduleService{}

	alice.CreateContainer(services, pdiService, scheduleService)

	return services, nil

}
