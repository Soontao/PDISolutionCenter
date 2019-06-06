package routes

import (
	"github.com/Soontao/PDISolutionCenter/modules/pdiclients"
	"github.com/Soontao/PDISolutionCenter/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// WithRoutes func
func WithRoutes(e *gin.Engine, db *gorm.DB, clients *pdiclients.PDIClients, ss *services.Services) {

	Wrapper := func(e ContextEndpointFunc) gin.HandlerFunc {
		return WithContext(db, clients, e, ss)
	}

	// api root
	api := e.Group("/api/v1")

	// user group
	userGroup := api.Group("/user")
	userGroup.GET("/", Wrapper(GetUserInformationEndpoint))

	// tenant group
	tenantGroup := api.Group("/tenant")
	tenantGroup.GET("/", Wrapper(GetCurrentUserTenant))
	tenantGroup.GET("/:id", Wrapper(GetTenantDetails))
	tenantGroup.PUT("/", Wrapper(AddNewTenant))

}
