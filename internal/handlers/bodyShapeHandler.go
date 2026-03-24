package handlers

import (
	"guitar-api/internal/services"
	"net/http"
	"github.com/gin-gonic/gin"
)

type BodyShapeHandler struct {
	Service *services.BodyShapeService
}

func (handler *BodyShapeHandler) GetBodyShapes(ctx *gin.Context) {
	bodyShapes, err := handler.Service.GetAllBodyShape()
	if err != nil  {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, bodyShapes)
}