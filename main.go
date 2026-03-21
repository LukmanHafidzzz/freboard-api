package main

import (
	"log"
	"guitar-api/internal/config"
	
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	log.Println("Starting app...")

	config.LoadEnv()
	config.NewDB()

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	log.Println("Server running on port 3000")
	router.Run(":3000")
}
