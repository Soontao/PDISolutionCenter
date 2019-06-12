package services

import (
	"github.com/Soontao/PDISolutionCenter/models"
	"github.com/jinzhu/gorm"
	"gopkg.in/robfig/cron.v2"
)

// ScheduleService type
type ScheduleService struct {
	db       *gorm.DB
	cron     *cron.Cron
	services *Services
}

// init service
func (s ScheduleService) init() {
	jobs := []*models.Schedule{}

	s.db.Find(jobs)

	// recover jobs from DB
	for _, job := range jobs {
		jobRunTimeID, _ := s.cron.AddJob(job.PeriodCron, job)
		job.ServerRuntimeID = int(jobRunTimeID)
		s.db.Save(job)
	}
}

// CreateRunningFunc func
func (s ScheduleService) CreateRunningFunc(job *models.Schedule) func() {
	return func() {

	}
}

// AddJob func
func (s ScheduleService) AddJob(job *models.Schedule) (*models.Schedule, error) {
	if err := s.db.Save(job).Error; err != nil {
		return nil, err
	}
	if jobRunningID, err := s.cron.AddFunc(job.PeriodCron, s.CreateRunningFunc(job)); err != nil {
		return nil, err
	} else {
		job.ServerRuntimeID = int(jobRunningID)
		s.db.Save(job)
	}
	return job, nil
}
