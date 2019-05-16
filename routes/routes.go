package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// WithRoutes func
func WithRoutes(e *gin.Engine, db *gorm.DB) {

	// userGroup
	userGroup := e.Group("/user")
	userGroup.GET("/", WithDatabaseAndSession(db, UserEndpoint))

}
