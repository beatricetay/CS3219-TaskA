package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "bye from web B!",
		})
	})
	r.Run(":8000") // listen and serve on port 8000
}
