package main

import (
	"github.com/Soontao/PDISolution/models"
	"github.com/Soontao/PDISolution/modules/memsession"
	"github.com/Soontao/PDISolution/modules/oauth"
	"github.com/Soontao/PDISolution/routes"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

// RunServer func
func RunServer(c *cli.Context) (err error) {

	db, err := models.CreateDB(c.GlobalString("db_type"), c.GlobalString("db_conn"))

	if err != nil {
		return err
	}
	// new server instance
	r := gin.Default()

	// with in memory session
	memsession.WithMemSession(r)

	// oauth config
	oConfig := &oauth.ServerOAuthConfig{
		AuthURL:      c.GlobalString("oauth_auth_url"),
		TokenURL:     c.GlobalString("oauth_token_url"),
		UserAPI:      c.GlobalString("oauth_user_api"),
		ClientID:     c.GlobalString("oauth_client_id"),
		ClientSecret: c.GlobalString("oauth_client_secret"),
		CallbackPath: "/oauth/callback",
		OnUserReceived: func(c *gin.Context, u oauth.SSOUser) (rt error) {
			session := sessions.Default(c)
			// set fed id
			session.Set(oauth.KeyFedID, u.GetFederationID())
			session.Save()
			return rt
		},
		OnCheckUser: func(c *gin.Context) (rt bool) {
			session := sessions.Default(c)
			if session.Get(oauth.KeyFedID) == nil {
				rt = false
			} else {
				// do more check
				rt = true
			}
			return rt
		},
	}

	// with oauth config
	oauth.WithOAuth(r, oConfig)

	// mount all routes
	routes.WithRoutes(r, db)

	// start server
	err = r.Run("0.0.0.0:18080")

	return err
}
