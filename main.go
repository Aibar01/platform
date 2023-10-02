package main

import (
	"github.com/Aibar01/platform/middleware/pfm"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	app := gin.New()

	app.Use(pfm.New())

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	app.Run(":6000")
}
