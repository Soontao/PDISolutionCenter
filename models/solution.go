package models

// Solution type
type Solution struct {
	BaseModel
	CurrentVersion     int
	CurrentStatus      string
	RecentCheckMessage string `gorm:"size:10240"`
	Contact            string
	ContactEmail       string
	PatchSolution      bool
	Tenant             *Tenant `gorm:"many2many:tenant_solution_rel;"`
}
