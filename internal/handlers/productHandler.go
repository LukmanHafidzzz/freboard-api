package handlers

import (
	"guitar-api/internal/models"
	"guitar-api/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Service *services.ProductService
}

func (handler *ProductHandler) GetProducts(ctx *gin.Context) {
	products, err := handler.Service.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	if products == nil {
		products = []models.Product{}
	}
	ctx.JSON(http.StatusOK, gin.H{"data": products})
}

func (handler *ProductHandler) GetProductById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	product, err := handler.Service.GetProductById(id)
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