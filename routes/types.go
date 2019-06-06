package routes

import (
	"github.com/Soontao/PDISolutionCenter/services"
	"github.com/Soontao/PDISolutionCenter/models"
	"github.com/Soontao/PDISolutionCenter/modules/oauth"
	"github.com/Soontao/PDISolutionCenter/modules/pdiclients"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// PlainJSONObject type
type PlainJSONObject map[string]interface{}

// ContextEndpointFunc func
type ContextEndpointFunc func(ctx *RouteContext)

// RouteContext type
type RouteContext struct {
	// HTTP Context
	HTTP *gin.Context
	// Database Context
	DB *gorm.DB
	// Session
	Session sessions.Session
	// PDI Clients
	PDI *pdiclients.PDIClients
	// Services
	Services *services.Services
}

// WithContext func
func WithContext(db *gorm.DB, pdi *pdiclients.PDIClients, endpoint ContextEndpointFunc, ss *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &RouteContext{
			HTTP:    c,
			DB:      db,
			Session: sessions.Default(c),
			PDI:     pdi,
			Services:ss,
		}
		endpoint(ctx)
	}
}

// GetCurrentUser information
func (c *RouteContext) GetCurrentUser() ( *models.User) {
	ssoID := c.Session.Get(oauth.KeyFedID).(string)
	rt := &models.User{}
	c.DB.Where(&models.User{FederationLoginID: ssoID}).First(rt)
	return rt
}
