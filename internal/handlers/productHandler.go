package handlers

import (
	"guitar-api/internal/models"
	"guitar-api/internal/services"
	"net/http"

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
