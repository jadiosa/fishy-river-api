package main

import (
	"github.com/jadiosa/fishy-river-api/Godeps/_workspace/src/github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080") // listen and server on 0.0.0.0:8080
}
