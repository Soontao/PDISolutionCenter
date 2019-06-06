package services

import (
	"github.com/Soontao/PDISolutionCenter/models"
	"github.com/jinzhu/gorm"
	"github.com/magic003/alice"
	"gopkg.in/robfig/cron.v2"
)

// ScheduleService type
type ScheduleService struct {
	alice.BaseModule
	DB   *gorm.DB
	cron *cron.Cron
}

// init service
func (s ScheduleService) init() {
	jobs := []*models.Schedule{}

	s.DB.Find(jobs)

	// recover jobs from DB
	for _, job := range jobs {
		jobRunTimeID, _ := s.cron.AddJob(job.PeriodCron, job)
		job.ServerRuntimeID = int(jobRunTimeID)
		s.DB.Save(job)
	}
}

// CreateRunningFunc func
func (s ScheduleService) CreateRunningFunc(job *models.Schedule) func() {
	return func() {

	}
}

// AddJob func
func (s ScheduleService) AddJob(job *models.Schedule) (*models.Schedule, error) {
	if err := s.DB.Save(job).Error; err != nil {
		return nil, err
	}
	if jobRunningID, err := s.cron.AddFunc(job.PeriodCron, s.CreateRunningFunc(job)); err != nil {
		return nil, err
	} else {
		job.ServerRuntimeID = int(jobRunningID)
		s.DB.Save(job)
	}
	return job, nil
}

// NewScheduleService service
func NewScheduleService(db *gorm.DB) *ScheduleService {

	ss := &ScheduleService{
		DB:   db,
		cron: cron.New(),
	}

	return ss

}
