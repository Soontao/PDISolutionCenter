package services

import (
	"errors"
	"fmt"
	"time"

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
	// do nothing
}

// newJobRunning object
func newJobRunning(db *gorm.DB, job *models.Schedule) *models.JobRunLog {
	runLog := &models.JobRunLog{
		Running: true,
	}
	db.Model(job).Association("RunInstances").Append(runLog)
	return runLog
}

// appendLog
func appendLog(db *gorm.DB, jobRun *models.JobRunLog, logLevel models.LogLevel, message string) {
	logItem := &models.Log{
		Level:   string(logLevel),
		Message: message,
	}
	db.Model(jobRun).Association("LogItems").Append(logItem)
}

func setJobRunSuccessFinished(db *gorm.DB, jobRun *models.JobRunLog) {
	jobRun.Successful = true
	jobRun.Running = false
	db.Save(jobRun)
}

func setJobRunFailed(db *gorm.DB, jobRun *models.JobRunLog) {
	jobRun.Successful = false
	jobRun.Running = false
	db.Save(jobRun)
}

func setJobRunTotalFailed(db *gorm.DB, jobRun *models.JobRunLog, err error) {
	appendLog(db, jobRun, models.LogLevelError, err.Error())
	setJobRunFailed(db, jobRun)
}

// CreateRunningFunc func
func (s ScheduleService) CreateRunningFunc(job *models.Schedule) func() {

	return func() {

		jobRun := newJobRunning(s.db, job)

		logInfo := func(message string) {
			appendLog(s.db, jobRun, models.LogLevelInfo, message)
		}

		logError := func(message string) {
			appendLog(s.db, jobRun, models.LogLevelError, message)
		}

		setJobFailed := func(message string) {
			setJobRunTotalFailed(s.db, jobRun, errors.New(message))
		}

		// on running job

		sourceTenant, err := s.services.PDIService.GetPDIClient(job.Solution.Tenant)

		if err != nil {
			setJobFailed(err.Error())
			return
		}

		sourceSolutionName := job.Solution.Name

		switch models.ScheduleType(job.ScheduleType) {

		case models.ScheduleTypeBuildCheck:

			logInfo("start check errors")

			buildErrors := sourceTenant.CheckBuildErrors(sourceSolutionName)

			if len(buildErrors) > 0 {

				for _, e := range buildErrors {
					logError(fmt.Sprintf("[%v]: %v", e.Catalog, e.Message))
				}

				setJobFailed("Errors found in solution")

				return
			}

		case models.ScheduleTypeBuildOnly:

			logInfo("start check errors")

			buildErrors := sourceTenant.CheckBuildErrors(sourceSolutionName)

			if len(buildErrors) > 0 {

				for _, e := range buildErrors {
					logError(fmt.Sprintf("[%v]: %v", e.Catalog, e.Message))
				}

				setJobFailed("Errors found in solution")

				return
			}

			logInfo("check finished, all things seems well")

			logInfo("start activation")

			if err = sourceTenant.ActivationSolution(sourceSolutionName); err != nil {
				setJobFailed(err.Error())
				return
			}

			logInfo("activation finished")

			logInfo("start assemble")

			if err = sourceTenant.AssembleSolution(sourceSolutionName); err != nil {
				setJobRunTotalFailed(s.db, jobRun, err)
				return
			}

			logInfo("assemble finished")

			logInfo("start create patch solution")

			if err = sourceTenant.CreatePatch(sourceSolutionName); err != nil {
				setJobRunTotalFailed(s.db, jobRun, err)
				return
			}

			logInfo("create patch finished")

		case models.ScheduleTypeDeployOnly:

		case models.ScheduleTypeBuildDeploy:

		default:
			setJobFailed(fmt.Sprintf("Not supported job type: %v", job.ScheduleType))
		}

		setJobRunSuccessFinished(s.db, jobRun)

	}
}

// AddJob func
func (s ScheduleService) AddJob(job *models.Schedule) (*models.Schedule, error) {

	if err := s.db.Save(job).Error; err != nil {
		return nil, err
	}

	go func() {
		// wait
		time.Sleep(time.Until(*job.OneTimeScheduleDateTime))
		// run
		s.CreateRunningFunc(job)()

	}()

	s.db.Save(job)

	return job, nil
}

func (s ScheduleService) DeleteJob(job *models.Schedule) (*models.Schedule, error) {

	return job, nil
}

func (s ScheduleService) DisableJob(job *models.Schedule) (*models.Schedule, error) {

	return job, nil
}

func (s ScheduleService) EnableJob(job *models.Schedule) (*models.Schedule, error) {

	return job, nil
}
