package service

import "github.com/dwiprastyoisworo/go-grpc-hexagonal/internal/product/port"

type ProductService struct {
	ProductRepository port.ProductRepoInterface
}

func NewProductService(productRepository port.ProductRepoInterface) port.ProductServiceInterface {
	return &ProductService{ProductRepository: productRepository}
}
