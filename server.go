package main

import (
	"github.com/Soontao/PDISolutionCenter/models"
	"github.com/Soontao/PDISolutionCenter/modules/memsession"
	"github.com/Soontao/PDISolutionCenter/modules/oauth"
	"github.com/Soontao/PDISolutionCenter/modules/pdiclients"
	"github.com/Soontao/PDISolutionCenter/routes"
	"github.com/Soontao/PDISolutionCenter/services"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

// RunServer func
func RunServer(c *cli.Context) (err error) {

	clients := pdiclients.NewPDIClients()

	db, err := models.CreateDB(c.GlobalString("db_type"), c.GlobalString("db_conn"))

	db.LogMode(c.GlobalBool("show_sql"))

	if err != nil {
		return err
	}

	ss, err := services.NewServices(db)

	if err != nil {
		return err
	}

	// new server instance
	r := gin.Default()

	// with in memory session
	memsession.WithMemSession(r)

	// OAuth config
	oConfig := &oauth.ServerOAuthConfig{
		AuthURL:      c.GlobalString("oauth_auth_url"),
		TokenURL:     c.GlobalString("oauth_token_url"),
		UserAPI:      c.GlobalString("oauth_user_api"),
		ClientID:     c.GlobalString("oauth_client_id"),
		ClientSecret: c.GlobalString("oauth_client_secret"),
		CallbackPath: "/oauth/callback",
		OnUserReceived: func(c *gin.Context, u oauth.SSOUser) (rt error) {

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
	routes.WithRoutes(r, db, clients, ss)

	// with static sources
	r.Use(static.Serve("/", static.LocalFile(c.GlobalString("static_path"), true)))

	// start server
	err = r.Run("0.0.0.0:18080")

	return err
}
