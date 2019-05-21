package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// WithRoutes func
func WithRoutes(e *gin.Engine, db *gorm.DB) {

	// api root
	api := e.Group("/api/v1")

	// userGroup
	userGroup := api.Group("/user")
	userGroup.GET("/", WithDatabaseAndSession(db, UserEndpoint))

}
