package app

import "github.com/gin-gonic/gin"

func ConfigureMappings(router *gin.Engine) *gin.Engine {
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	return router
}
