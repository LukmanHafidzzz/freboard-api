package handlers

import (
	"guitar-api/internal/services"
	"net/http"
	"github.com/gin-gonic/gin"
)

type BrandHandler struct {
    Service *services.BrandService
}


func (handler *BrandHandler) GetBrands(ctx *gin.Context) {
	brands, err := handler.Service.GetAllBrands()
	if err != nil  {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, brands)
}