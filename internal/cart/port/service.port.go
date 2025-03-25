package port

import (
	"context"
	"github.com/dwiprastyoisworo/go-grpc-hexagonal/internal/cart/model"
)

type CartServiceInterface interface {
	CreateCart(ctx context.Context, cart *model.CreateCartReq) error
	GetCart(ctx context.Context, userId string) ([]model.CartDto, error)
	DeleteCart(ctx context.Context, userId string, productId string) error
}
