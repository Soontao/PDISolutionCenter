package oauth

import (
	"fmt"

	"github.com/gin-contrib/sessions/memstore"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
	"golang.org/x/oauth2"
)

// WithOAuth config
func WithOAuth(e *gin.Engine) (rt error) {

	endpoint := oauth2.Endpoint{
		AuthURL:  "https://github.tools.sap/login/oauth/authorize",
		TokenURL: "https://github.tools.sap/login/oauth/access_token",
	}

	userAPI := "https://github.tools.sap/api/v3/user"

	oauth2Config := &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "http://localhost:18080/oauth/callback",
		Scopes:       []string{},
		Endpoint:     endpoint,
	}

	store := memstore.NewStore([]byte("secret"))

	e.Use(sessions.Sessions("session", store))

	e.GET("/oauth/callback", func(c *gin.Context) {

		code := c.Query("code")

		token, err := oauth2Config.Exchange(oauth2.NoContext, code)

		if err != nil {
			c.AbortWithError(500, err)
		}

		res, err := req.Get(userAPI, req.Header{"Authorization": fmt.Sprintf("token %s", token.AccessToken)})

		if err != nil {
			c.AbortWithError(500, err)
		}

		uString, _ := res.ToString()

		session := sessions.Default(c)
		session.Set("user", uString)
		session.Save()

		c.Redirect(302, "/")

	})

	e.Use(func(c *gin.Context) {

		if sessions.Default(c).Get("user") == nil {
			c.Redirect(302, oauth2Config.AuthCodeURL("0000"))
		} else {
			c.Next()
		}

	})

	return rt
}
