package routes

import (
	"strconv"

	"github.com/jinzhu/gorm"

	"github.com/Soontao/PDISolutionCenter/models"
)

// GetCurrentUserTenant endpoint
func GetCurrentUserTenant(c *RouteContext) {

	user := c.GetCurrentUser()

	c.DB.Preload("Tenants").Where(&models.User{FederationLoginID: user.FederationLoginID}).Find(user)

	c.HTTP.JSON(200, PlainJSONObject{"Tenants": user.Tenants})

}

type addTenantRequest struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Hostname    string `json:"Hostname" validate:"required"`
	Username    string `json:"Username" validate:"required"`
	Password    string `json:"Password" validate:"required"`
}

// AddNewTenant endpoint
func AddNewTenant(c *RouteContext) {
	currentUser := c.GetCurrentUser()

	payload := &addTenantRequest{}

	c.HTTP.Bind(payload)

	// should check another people have added it.
	pdi, err := c.PDI.GetOrAddNewClient(payload.Hostname, payload.Username, payload.Password)

	if err != nil {
		c.HTTP.AbortWithStatusJSON(500, PlainJSONObject{"error": err.Error()})
		return
	}

	tx := c.DB.Begin()

	// add refresh jobs

	t := &models.Tenant{
		BaseModel: models.BaseModel{
			Name:        payload.Name,
			Description: payload.Description,
			UpdatedBy:   currentUser,
			CreatedBy:   currentUser,
		},
		TenantHost:         payload.Hostname,
		TenantUser:         payload.Username,
		TenantUserPassword: payload.Password,
	}

	tx.Save(t)

	// save solution list

	for _, s := range pdi.GetSolutionsAPI() {

		status := pdi.GetSolutionStatus(s.Name)

		s := &models.Solution{
			BaseModel: models.BaseModel{
				Name:        s.Name,
				Description: s.Description,
				UpdatedBy:   currentUser,
				CreatedBy:   currentUser,
			},
			CurrentVersion: int(status.Version),
			PatchSolution:  s.PatchSolution,
			Contact:        s.Contact,
			ContactEmail:   s.Email,
			CurrentStatus:  s.Status,
		}

		tx.Create(s)

		tx.Model(t).Association("Solutions").Append(s)

	}

	tx.Model(t).Association("Admins").Append(currentUser)

	tx.Commit()

	c.HTTP.JSON(201, PlainJSONObject{"successful": true})

}

// GetTenantDetails endpoint
func GetTenantDetails(c *RouteContext) {

	tenant := &models.Tenant{}

	if tenantID, err := strconv.ParseUint(c.HTTP.Param("id"), 10, 32); err != nil {
		c.HTTP.AbortWithStatusJSON(500, PlainJSONObject{"error": err.Error()})
		return
	} else {
		if err = c.DB.Preload("Solutions").Preload("Admins").First(tenant, tenantID).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				c.HTTP.AbortWithStatusJSON(404, PlainJSONObject{"error": err.Error()})
			} else {
				c.HTTP.AbortWithStatusJSON(500, PlainJSONObject{"error": err.Error()})
			}
			return
		}
		c.HTTP.JSON(200, tenant)
		return
	}

}
