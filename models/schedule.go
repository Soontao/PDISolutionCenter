package models

// ScheduleType enum
type ScheduleType string

const (
	// ScheduleTypeBuildCheck enum value
	ScheduleTypeBuildCheck = ScheduleType("buildcheck")
	// ScheduleTypeDeployOnly enum value
	ScheduleTypeDeployOnly = ScheduleType("deploy")
	// ScheduleTypeBuildOnly enum value
	ScheduleTypeBuildOnly = ScheduleType("buildonly")
	// ScheduleTypeBuildDeploy enum value
	ScheduleTypeBuildDeploy = ScheduleType("builddeploy")
)

// Schedule type
type Schedule struct {
	BaseModel
	ScheduleType    string
	PeriodCron      string
	Solution        *Solution
	TargetTenant    *Tenant
	RunInstances    []*JobRunLog
	Enable          bool
	ServerRuntimeID int
}

// GetScheduleType enum
func (s *Schedule) GetScheduleType() ScheduleType {
	return ScheduleType(s.ScheduleType)
}
