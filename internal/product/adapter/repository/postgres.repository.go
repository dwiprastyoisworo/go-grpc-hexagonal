package repository

import (
	"context"
	"github.com/dwiprastyoisworo/go-grpc-hexagonal/internal/product/model"
	"log"
)

func (p *ProductRepository) GetAll(ctx context.Context) ([]model.ProductDto, error) {
	var products []model.ProductDto
	err := p.db.SelectContext(ctx, &products, "SELECT id, name, price, image_url, quantity FROM products")
	return products, err
}

func (p *ProductRepository) Create(ctx context.Context, product *model.ProductDto) error {
	stmt, err := p.db.PrepareNamedContext(ctx, "INSERT INTO products (id, name, price, image_url, quantity) VALUES (:id, :name, :price, :image_url, :quantity)")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, product)
	return err
}

func (p *ProductRepository) UpdateQuantity(ctx context.Context, qty int, id string) error {
	stmt, err := p.db.PrepareNamedContext(ctx, "UPDATE products SET quantity = :quantity WHERE id = :id")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, map[string]interface{}{"quantity": qty, "id": id})
	return err
}
