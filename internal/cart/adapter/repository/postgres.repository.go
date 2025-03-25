package repository

import (
	"context"
	"github.com/dwiprastyoisworo/go-grpc-hexagonal/internal/cart/model"
	"log"
)

func (c *CartRepository) CreateCart(ctx context.Context, cart *model.CartDto) error {
	stmt, err := c.db.PrepareNamedContext(ctx, "INSERT INTO carts (id, user_id, product_id, quantity) VALUES (:id, :user_id, :product_id, :quantity)")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, cart)
	return err
}

func (c *CartRepository) GetCart(ctx context.Context, userId string) ([]model.CartDto, error) {
	var carts []model.CartDto
	err := c.db.SelectContext(ctx, &carts, "SELECT id, user_id, product_id, quantity FROM carts WHERE user_id = $1", userId)
	return carts, err
}

func (c *CartRepository) DeleteCart(ctx context.Context, userId string, productId string) error {
	stmt, err := c.db.PrepareNamedContext(ctx, "DELETE FROM carts WHERE user_id = :user_id AND product_id = :product_id")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, map[string]interface{}{"user_id": userId, "product_id": productId})
	return err
}
