package v1

import (
	"guitar-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupBrandRoutes(router *gin.RouterGroup, handler *handlers.BrandHandler) {
	brandRoutes := router.Group("/brands")
	{
		brandRoutes.GET("", handler.GetBrands)
		brandRoutes.GET("/:id", handler.GetBrandById)
	}
}