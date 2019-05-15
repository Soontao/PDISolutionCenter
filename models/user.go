package models

// User type
type User struct {
	BaseModel
	FederationLoginID string `gorm:"unique"`
	Email             string
	Tenants           []*Tenant `gorm:"many2many:tenant_user_rel;"`
}
