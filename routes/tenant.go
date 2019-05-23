package routes

import (
	"github.com/Soontao/PDISolutionCenter/models"
)

// GetCurrentUserTenant endpoint
func GetCurrentUserTenant(c *RouteContext) {
	user := c.GetCurrentUser()
	c.DB.Model(user).Association("Tenants")
	c.HTTP.JSON(200, PlainJSONObject{"Tenants": user.Tenants})
}

type addTenantRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Hostname    string `json:"hostname" validate:"required"`
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

// AddNewTenant endpoint
func AddNewTenant(c *RouteContext) {
	currentUser := c.GetCurrentUser()

	payload := &addTenantRequest{}

	c.HTTP.Bind(payload)

	// should check another people have added it.
	pdi, err := c.PDI.GetOrAddNewClient(payload.Hostname, payload.Username, payload.Password)

	if err != nil {
		c.HTTP.Error(err)
	}

	// save solution list

	solutions := []*models.Solution{}

	for _, s := range pdi.GetSolutionsAPI() {
		solutions = append(solutions, &models.Solution{
			BaseModel: models.BaseModel{
				Name:        s.Name,
				Description: s.Description,
				UpdatedBy:   currentUser,
				CreatedBy:   currentUser,
			},
			Contact:       s.Contact,
			ContactEmail:  s.Email,
			CurrentStatus: s.Status,
		})
	}

	// add refresh jobs

	t := models.Tenant{
		BaseModel: models.BaseModel{
			Name:        payload.Name,
			Description: payload.Description,
			UpdatedBy:   currentUser,
			CreatedBy:   currentUser,
		},
		TenantHost:         payload.Hostname,
		TenantUser:         payload.Username,
		TenantUserPassword: payload.Password,
		Solutions:          solutions,
		Admins:             []*models.User{currentUser},
	}

	c.DB.Save(t)

	c.HTTP.JSON(200, PlainJSONObject{"successful": true})
}
