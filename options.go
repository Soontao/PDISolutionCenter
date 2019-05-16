package main

import "github.com/urfave/cli"

var options = []cli.Flag{
	cli.StringFlag{
		Name:   "oauth_auth_url",
		EnvVar: "OAUTH_AUTH_URL",
	},
	cli.StringFlag{
		Name:   "oauth_token_url",
		EnvVar: "OAUTH_TOKEN_URL",
	},
	cli.StringFlag{
		Name:   "oauth_user_api",
		EnvVar: "OAUTH_USER_API",
	},
	cli.StringFlag{
		Name:   "oauth_client_id",
		EnvVar: "OAUTH_CLIENT_ID",
	}, cli.StringFlag{
		Name:   "oauth_client_secret",
		EnvVar: "OAUTH_CLIENT_SECRET",
	},
}
