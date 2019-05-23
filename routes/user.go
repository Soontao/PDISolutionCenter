package routes

import (
	"github.com/Soontao/PDISolutionCenter/modules/oauth"
)

// UserEndpoint route
func UserEndpoint(c *RouteContext) {
	userID := c.Session.Get(oauth.KeyFedID)
	c.HTTP.JSON(200, map[string]interface{}{"User": userID})
}
