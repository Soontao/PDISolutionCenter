// Package services apis
// PDI Package Services
package services

import (
	"github.com/Soontao/PDISolutionCenter/models"
	pdiutil "github.com/Soontao/pdi-util"
	"github.com/jinzhu/gorm"
)

// PDIService type
type PDIService struct {
	db       *gorm.DB
	services *Services
}

// GetPDIClient object
func (p PDIService) GetPDIClient(t *models.Tenant) (*pdiutil.PDIClient, error) {
	return pdiutil.NewPDIClient(t.TenantUser, t.TenantUserPassword, t.TenantHost, "")
}
