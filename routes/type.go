package routes

import (
	"github.com/Soontao/PDISolutionCenter/modules/pdiclients"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// RouteContext type
type RouteContext struct {
	// HTTP Context
	HTTP    *gin.Context
	// Database Context
	DB      *gorm.DB
	// Session
	Session sessions.Session
	// PDI Clients
	PDI     *pdiclients.PDIClients
}

// EndpointFunc func
type EndpointFunc func(*gin.Context, *gorm.DB, sessions.Session)

// ContextEndpointFunc func
type ContextEndpointFunc func(ctx *RouteContext)

// WithContext func
func WithContext(db *gorm.DB, pdi *pdiclients.PDIClients, endpoint ContextEndpointFunc) (rt func(*gin.Context)) {
	return func(c *gin.Context) {
		ctx := &RouteContext{
			HTTP:    c,
			DB:      db,
			Session: sessions.Default(c),
			PDI:     pdi,
		}
		endpoint(ctx)
	}
}

// WithDatabaseAndSession func
func WithDatabaseAndSession(db *gorm.DB, endpoint EndpointFunc) (rt func(*gin.Context)) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		endpoint(c, db, session)
	}
}
