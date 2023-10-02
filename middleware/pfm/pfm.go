package pfm

import (
	"github.com/Aibar01/platform/platform"
	"github.com/gin-gonic/gin"
	"net/http"
)

func New() gin.HandlerFunc {
	return func(c *gin.Context) {
		pfm := &platform.PlatformContext{}
		if err := pfm.Validate(c.Request.Header); err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, map[string]string{"error": err.Error()})
			return
		}
		c.Set("pfm", pfm)
	}
}
