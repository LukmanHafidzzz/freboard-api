package main

import (
	"context"
	v1 "guitar-api/api/v1"
	"guitar-api/internal/config"
	"guitar-api/internal/handlers"
	middleware "guitar-api/internal/middlewares"
	"guitar-api/internal/services"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting app...")

	config.LoadEnv()
	config.NewDB()

	brandService := &services.BrandService{DB: config.DB}
	bodyShapeService := &services.BodyShapeService{DB: config.DB}
	productService := &services.ProductService{DB: config.DB}

	brandHandler := &handlers.BrandHandler{Service: brandService}
	bodyShapeHandler := &handlers.BodyShapeHandler{Service: bodyShapeService}
	productHandler := &handlers.ProductHandler{Service: productService}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(middleware.CORS())
	router.Use(middleware.Gzip())
	router.Use(middleware.RateLimiter())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	v1Router := router.Group("/api/v1")
	v1.SetupBrandRoutes(v1Router, brandHandler)
	v1.SetupBodyShapeRoutes(v1Router, bodyShapeHandler)
	v1.SetupProductRoutes(v1Router, productHandler)

	srv := &http.Server{
		Addr:         ":" + config.GetEnv("PORT"),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Println("Server running on port " + config.GetEnv("PORT"))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server error: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exited")
}