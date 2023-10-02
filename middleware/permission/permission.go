package permission

import (
	"github.com/Aibar01/platform/response"
	"github.com/Aibar01/platform/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Copy()
		pfm := utils.GetPFM(c)
		if !pfm.User.IsAdmin() {
			c.AbortWithStatusJSON(http.StatusForbidden, response.PermissionDeniedError)
			return
		}

		c.Next()
	}
}

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		pfm := utils.GetPFM(c)
		if !pfm.User.IsAdmin() {
			c.AbortWithStatusJSON(http.StatusForbidden, response.PermissionDeniedError)
			return
		}

		c.Next()
	}
}

func SimplePermission(permissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		pfm := utils.GetPFM(c)
		for _, permission := range permissions {
			if !pfm.User.HasPermission(permission) {
				c.AbortWithStatusJSON(http.StatusForbidden, response.PermissionDeniedError)
				return
			}
		}

		c.Next()
	}
}

func IsAdminOrSimplePermission(permissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		pfm := utils.GetPFM(c)
		if !pfm.User.IsAdmin() {
			c.AbortWithStatusJSON(http.StatusForbidden, response.PermissionDeniedError)
			return
		}

		for _, permission := range permissions {
			if !pfm.User.HasPermission(permission) {
				c.AbortWithStatusJSON(http.StatusForbidden, response.PermissionDeniedError)
				return
			}
		}

		c.Next()
	}
}

func ConsumerPermission(consumerName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		pfm := utils.GetPFM(c)
		if pfm.Consumer.Name != consumerName {
			c.AbortWithStatusJSON(http.StatusForbidden, response.PermissionDeniedError)
			return
		}

		c.Next()
	}
}
