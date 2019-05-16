package routes

import (
	"github.com/Soontao/PDISolution/modules/oauth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// UserEndpoint route
func UserEndpoint(c *gin.Context, db *gorm.DB, s sessions.Session) {
	userID := s.Get(oauth.KeyFedID)
	c.JSON(200, map[string]interface{}{"User": userID})
}
