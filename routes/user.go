package routes

// GetUserInformationEndpoint route
//
// Get the basic user information of current user
func GetUserInformationEndpoint(c *RouteContext) {
	c.HTTP.JSON(200, map[string]interface{}{"User": c.GetCurrentUser()})
}
