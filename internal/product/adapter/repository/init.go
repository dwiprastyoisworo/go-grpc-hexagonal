package repository

import (
	"github.com/dwiprastyoisworo/go-grpc-hexagonal/internal/product/port"
	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) port.ProductRepoInterface {
	return &ProductRepository{db: db}
}
