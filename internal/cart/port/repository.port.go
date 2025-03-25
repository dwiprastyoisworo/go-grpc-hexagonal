package port

import (
	"context"
	"github.com/dwiprastyoisworo/go-grpc-hexagonal/internal/cart/model"
)

type CartRepositoryInterface interface {
	CreateCart(ctx context.Context, cart *model.CartDto) error
	GetCart(ctx context.Context, userId string) ([]model.CartDto, error)
	DeleteCart(ctx context.Context, userId string, productId string) error
}
