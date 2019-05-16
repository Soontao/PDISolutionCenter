package models

// Tenant type
type Tenant struct {
	BaseModel
	TenantHost           string `gorm:"unique"`
	TenantUser           string
	TenantUserPassword   string
	TenantReleaseVersion string
	Admins               []*User `gorm:"many2many:tenant_user_rel;"`
	Solutions            []*Solution
}
