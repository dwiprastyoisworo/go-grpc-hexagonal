package handler

import (
	"github.com/dwiprastyoisworo/go-grpc-hexagonal/internal/product/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (p *ProductHandler) GetAll(ctx *gin.Context) {
	products, err := p.ProductService.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
	return
}

func (p *ProductHandler) Create(ctx *gin.Context) {
	var productReq model.CreateProductReq
	if err := ctx.ShouldBindJSON(&productReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := p.ProductService.Create(ctx, &productReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "product created"})
	return
}
