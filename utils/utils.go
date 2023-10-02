package utils

import (
	"github.com/Aibar01/platform/platform"
	"github.com/gin-gonic/gin"
)

func GetPFM(c *gin.Context) *platform.PlatformContext {
	pfm, _ := c.Get("pfm")

	return pfm.(*platform.PlatformContext)
}
