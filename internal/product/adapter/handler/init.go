package handler

import "github.com/dwiprastyoisworo/go-grpc-hexagonal/internal/product/port"

type ProductHandler struct {
	ProductService port.ProductServiceInterface
}

func NewProductHandler(productService port.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{ProductService: productService}
}
