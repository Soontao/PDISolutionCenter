package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// WithDatabase func
func WithDatabase(db *gorm.DB, route func(*gin.Context, *gorm.DB)) (rt func(*gin.Context)) {
	return func(c *gin.Context) {
		route(c, db)
	}
}

// EndpointFunc func
type EndpointFunc func(*gin.Context, *gorm.DB, sessions.Session)
