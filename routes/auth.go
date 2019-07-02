package routes

import (
	"github.com/Soontao/PDISolutionCenter/models"
	"github.com/Soontao/PDISolutionCenter/modules/oauth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// CreateOAuthOnUserReceived func
func CreateOAuthOnUserReceived(db *gorm.DB) func(c *gin.Context, u oauth.SSOUser) (rt error) {
	return func(c *gin.Context, u oauth.SSOUser) (rt error) {

		session := sessions.Default(c)

		fedID := u.GetFederationID()

		// set fed id
		session.Set(oauth.KeyFedID, fedID)
		session.Save()

		// sync user to db
		user := &models.User{}
		query := &models.User{FederationLoginID: fedID}
		update := &models.User{
			Email:     u.GetEmail(),
			BaseModel: models.BaseModel{Name: u.GetUserName()},
		}

		rt = db.Where(query).Assign(update).FirstOrCreate(user).Error

		return rt
	}
}

// OAuthOnCheckUser func
func OAuthOnCheckUser(c *gin.Context) (rt bool) {
	session := sessions.Default(c)
	if session.Get(oauth.KeyFedID) == nil {
		rt = false
	} else {
		// do more check
		rt = true
	}
	return rt
}
