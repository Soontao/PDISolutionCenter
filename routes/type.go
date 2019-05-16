package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// EndpointFunc func
type EndpointFunc func(*gin.Context, *gorm.DB, sessions.Session)

// WithDatabaseAndSession func
func WithDatabaseAndSession(db *gorm.DB, endpoint EndpointFunc) (rt func(*gin.Context)) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		endpoint(c, db, session)
	}
}
