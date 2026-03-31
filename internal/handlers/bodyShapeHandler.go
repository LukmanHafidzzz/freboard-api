package handlers

import (
	"guitar-api/internal/services"
	"net/http"
	"strconv"

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
	ctx.JSON(http.StatusOK, gin.H{"data": bodyShapes})
}

func (handler *BodyShapeHandler) GetBodyShapeById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	bodyShape, err := handler.Service.GetBodyShapeById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if bodyShape == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "body shape not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": bodyShape})
}

func (handler *BodyShapeHandler) GetProductsByBodyShapeId(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	product, err := handler.Service.GetAllProductsByBodyShapeId(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if product == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": product})
}