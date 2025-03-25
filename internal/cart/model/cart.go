package model

type CartDto struct {
	Id        string `json:"id" db:"id"`
	UserId    string `json:"user_id" db:"user_id"`
	ProductId string `json:"product_id" db:"product_id"`
	Quantity  int    `json:"quantity" db:"quantity"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type CreateCartReq struct {
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
