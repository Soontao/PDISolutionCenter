package memsession

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

// WithMemSession func
func WithMemSession(c *gin.Engine) {

	store := memstore.NewStore([]byte(uuid.NewV4().String()))

	c.Use(sessions.Sessions("application", store))

}
