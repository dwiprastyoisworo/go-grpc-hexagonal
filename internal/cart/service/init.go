package service

import "github.com/dwiprastyoisworo/go-grpc-hexagonal/internal/cart/port"

type CartService struct {
	CartRepository port.CartRepositoryInterface
}

func NewCartService(cartRepository port.CartRepositoryInterface) *CartService {
	return &CartService{CartRepository: cartRepository}
}
