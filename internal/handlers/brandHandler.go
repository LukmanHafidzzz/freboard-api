package handlers

import (
	"guitar-api/internal/models"
	"guitar-api/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BrandHandler struct {
	Service *services.BrandService
}

func (handler *BrandHandler) GetBrands(ctx *gin.Context) {
	brands, err := handler.Service.GetAllBrands()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if brands == nil {
		brands = []models.Brand{}
	}
	ctx.JSON(http.StatusOK, gin.H{"data": brands})
}

func (handler *BrandHandler) GetBrandById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	brand, err := handler.Service.GetBrandById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if brand == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "brand not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": brand})
}
