package service

import (
	"context"
	"github.com/dwiprastyoisworo/go-grpc-hexagonal/internal/product/model"
	"github.com/google/uuid"
)

func (p *ProductService) GetAll(ctx context.Context) ([]model.ProductDto, error) {
	return p.ProductRepository.GetAll(ctx)
}

func (p *ProductService) Create(ctx context.Context, product *model.CreateProductReq) error {
	// TODO: add validation here

	productDto := model.ProductDto{
		ID:       uuid.NewString(),
		Name:     product.Name,
		Price:    product.Price,
		ImageUrl: product.ImageUrl,
		Quantity: product.Quantity,
	}
	return p.ProductRepository.Create(ctx, &productDto)
}

func (p *ProductService) UpdateQuantity(ctx context.Context, id string, quantity int) error {
	// TODO add validation here
	return p.ProductRepository.UpdateQuantity(ctx, quantity, id)
}
