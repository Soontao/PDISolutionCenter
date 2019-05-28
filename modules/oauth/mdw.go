package oauth

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
	"golang.org/x/oauth2"
)

// ServerOAuthConfig type
type ServerOAuthConfig struct {
	AuthURL        string
	TokenURL       string
	UserAPI        string
	ClientID       string
	ClientSecret   string
	CallbackPath   string
	OnUserReceived func(*gin.Context, SSOUser) error
	OnCheckUser    func(*gin.Context) bool
}

// GithubGetUserAPI func
func GithubGetUserAPI(userAPI, token string) (rt *GithubUser, err error) {

	res, err := req.Get(userAPI, req.Header{"Authorization": fmt.Sprintf("token %s", token)})
	if err != nil {
		return nil, err
	}
	rt = &GithubUser{}
	err = res.ToJSON(rt)
	if err != nil {
		return nil, err
	}
	return rt, err
}

// WithOAuth config
func WithOAuth(e *gin.Engine, config *ServerOAuthConfig) (rt error) {

	oauth2Config := &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Scopes:       []string{},
		Endpoint: oauth2.Endpoint{
			AuthURL:  config.AuthURL,
			TokenURL: config.TokenURL,
		},
	}

	e.GET(config.CallbackPath, func(c *gin.Context) {

		code := c.Query("code")

		state := c.Query("state")

		if code == "" {
			c.AbortWithError(400, fmt.Errorf("Bad OAuth Callback Request"))
		}

		if state == "" {
			state = "/"
		}

		token, err := oauth2Config.Exchange(oauth2.NoContext, code)

		if err != nil {
			c.AbortWithError(500, err)
		}

		user, err := GithubGetUserAPI(config.UserAPI, token.AccessToken)

		if err != nil {
			c.AbortWithError(500, err)
		}

		config.OnUserReceived(c, *user)

		c.Redirect(302, state)

	})

	e.Use(func(c *gin.Context) {

		if config.OnCheckUser(c) {
			c.Next()
		} else {
			// set state as original request uri
			c.Redirect(302, oauth2Config.AuthCodeURL(c.Request.RequestURI))
			c.Abort()
		}

	})

	return rt
}
