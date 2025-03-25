package service

import (
	"context"
	"github.com/dwiprastyoisworo/go-grpc-hexagonal/internal/cart/model"
)

func (c *CartService) CreateCart(ctx context.Context, cart *model.CreateCartReq) error {
	//TODO implement me
	panic("implement me")
}

func (c *CartService) GetCart(ctx context.Context, userId string) ([]model.CartDto, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CartService) DeleteCart(ctx context.Context, userId string, productId string) error {
	//TODO implement me
	panic("implement me")
}
