package main

import (
	"log"
	"os"

	"github.com/jadiosa/fishy-river-api/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	port := getPort();

	router := gin.Default()

	router.GET("/feed", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run(port) // listen and server on 0.0.0.0:8080
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
		log.Println("[-] No PORT environment variable detected. Setting to ", port)
	}
	return ":" + port
}
