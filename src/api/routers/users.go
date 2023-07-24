package routers

import (
	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup) {

	router.GET("/send-otp", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
}
