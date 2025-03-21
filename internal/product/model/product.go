package model

import "time"

type ProductDto struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Price     float64   `json:"price" db:"price"`
	ImageUrl  string    `json:"image_url" db:"image_url"`
	Quantity  int       `json:"quantity" db:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type CreateProductReq struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	ImageUrl string  `json:"image_url"`
	Quantity int     `json:"quantity"`
}
