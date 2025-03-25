package repository

import "github.com/jmoiron/sqlx"

type CartRepository struct {
	db *sqlx.DB
}

func NewCartRepository(db *sqlx.DB) *CartRepository {
	return &CartRepository{db: db}
}
