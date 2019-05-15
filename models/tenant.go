package models

// Tenant type
type Tenant struct {
	BaseModel
	TenantDevUser         string
	TenantDevUserPassword string
	TenantHost            string
	TenantReleaseVersion  string
	Admins                []*User `gorm:"many2many:tenant_user_rel;"`
	Solutions             []*Solution
}
