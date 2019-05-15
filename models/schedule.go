package models

import "time"

// Schedule type
type Schedule struct {
	BaseModel
	ScheduleType  string
	PeriodCron    string
	TargetRunTime *time.Time
	Solution      *Solution
	TargetTenant  *Tenant
	Logs          []*JobRunLog
}
