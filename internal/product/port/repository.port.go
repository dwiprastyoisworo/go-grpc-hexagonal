package port

import (
	"context"
	"github.com/dwiprastyoisworo/go-grpc-hexagonal/internal/product/model"
)

type ProductRepoInterface interface {
	GetAll(ctx context.Context) ([]model.ProductDto, error)
	Create(ctx context.Context, product *model.ProductDto) error
	UpdateQuantity(ctx context.Context, qty int, id string) error
}
