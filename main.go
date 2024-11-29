package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	router.GET("bye/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "bye bye world",
		})
	})

	router.Run(":5000")
}
