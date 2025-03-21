package port

import (
	"context"
	"github.com/dwiprastyoisworo/go-grpc-hexagonal/internal/product/model"
)

type ProductServiceInterface interface {
	GetAll(ctx context.Context) ([]model.ProductDto, error)
	Create(ctx context.Context, product *model.CreateProductReq) error
	UpdateQuantity(ctx context.Context, id string, quantity int) error
}
