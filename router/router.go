package router

import "github.com/gin-gonic/gin"

func Default(route *gin.Engine) error {
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return nil
}
