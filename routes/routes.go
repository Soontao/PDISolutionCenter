package routes

import (
	"github.com/Soontao/PDISolutionCenter/modules/pdiclients"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// WithRoutes func
func WithRoutes(e *gin.Engine, db *gorm.DB, clients *pdiclients.PDIClients) {

	// api root
	api := e.Group("/api/v1")

	// userGroup
	userGroup := api.Group("/user")
	userGroup.GET("/", WithContext(db,clients, UserEndpoint))

}
