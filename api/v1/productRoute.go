package v1

import (
	"guitar-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(router *gin.RouterGroup, handler *handlers.ProductHandler) {
	productRoutes := router.Group("/products")
	{
		productRoutes.GET("", handler.GetProducts)
	}
}