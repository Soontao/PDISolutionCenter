package main

import (
	"github.com/Soontao/PDISolution/modules/oauth"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	oauth.WithOAuth(r)

	r.Run("0.0.0.0:18080")

}
