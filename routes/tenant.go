package routes

import (
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

	// save solution list

	solutions := []*models.Solution{}

	for _, s := range pdi.GetSolutionsAPI() {
		s := &models.Solution{
			BaseModel: models.BaseModel{
				Name:        s.Name,
				Description: s.Description,
				UpdatedBy:   currentUser,
				CreatedBy:   currentUser,
			},
			Contact:       s.Contact,
			ContactEmail:  s.Email,
			CurrentStatus: s.Status,
		}
		c.DB.Create(s)
		solutions = append(solutions, s)
	}

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

	c.DB.Save(t)

	c.DB.Model(t).Association("Solutions").Append(solutions)

	c.DB.Model(t).Association("Admins").Append(currentUser)

	c.HTTP.JSON(201, PlainJSONObject{"successful": true})
}
