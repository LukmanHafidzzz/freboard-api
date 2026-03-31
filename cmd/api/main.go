package main

import (
	v1 "guitar-api/api/v1"
	"guitar-api/internal/config"
	"guitar-api/internal/handlers"
	"guitar-api/internal/services"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting app...")

	config.LoadEnv()
	config.NewDB()

	brandService := &services.BrandService{
		DB: config.DB,
	}
	bodyShapeService := &services.BodyShapeService{
		DB: config.DB,
	}

	brandHandler := &handlers.BrandHandler{
		Service: brandService,
	}

	bodyShapeHandler := &handlers.BodyShapeHandler{
		Service: bodyShapeService,
	}

	productService := &services.ProductService{
		DB: config.DB,
	}

	productHandler := &handlers.ProductHandler{
		Service: productService,
	}

	router := gin.Default()
	
	v1Router := router.Group("/api/v1")
	v1.SetupBrandRoutes(v1Router, brandHandler)
	v1.SetupBodyShapeRoutes(v1Router, bodyShapeHandler)
	v1.SetupProductRoutes(v1Router, productHandler)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	log.Println("Server running on port 3000")
	router.Run(":3000")
}