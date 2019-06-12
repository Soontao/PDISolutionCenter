package services

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/robfig/cron.v2"
)

// Service interface
type Service interface {
	init()
}

// Services type
type Services struct {
	PDIService      *PDIService
	ScheduleService *ScheduleService
}

// NewServices constructor
func NewServices(db *gorm.DB) (*Services, error) {

	services := &Services{}

	pdiService := &PDIService{
		db:       db,
		services: services,
	}

	scheduleService := &ScheduleService{
		db:       db,
		cron:     cron.New(),
		services: services,
	}

	services.PDIService = pdiService
	services.ScheduleService = scheduleService

	return services, nil

}
