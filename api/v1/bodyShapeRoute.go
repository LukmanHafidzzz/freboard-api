package v1

import (
	"guitar-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupBodyShapeRoutes(router *gin.RouterGroup, handler *handlers.BodyShapeHandler) {
	bodyShapeRoutes := router.Group("/body-shapes")
	{
		bodyShapeRoutes.GET("", handler.GetBodyShapes)
		bodyShapeRoutes.GET("/:id", handler.GetBodyShapeById)
	}
}