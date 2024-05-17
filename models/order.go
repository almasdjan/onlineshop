package models

type OrderForUser struct {
	Id          int     `json:"-" db:"id"`
	TotalAmount float64 `json:"total_amount" db:"total_amount"`
	Status      string  `json:"status"`
	CreatedAt   string  `json:"created_at" db:"created_date"`
	Address     string  `json:"address"`
	ProductId   int     `json:"product_id" db:"product_id"`
	Name        string  `json:"name"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

type OrderWithProducts struct {
	Id          int                   `json:"id" db:"id"`
	TotalAmount float64               `json:"total_amount" db:"total_amount"`
	Status      string                `json:"status"`
	CreatedAt   string                `json:"created_at" db:"created_date"`
	Address     string                `json:"address"`
	UserId      int                   `json:"user_id" db:"user_id"`
	Products    []GetProductsFromCart `json:"products"`
}

type OrderForAdmin struct {
	Id          int     `json:"-" db:"id"`
	TotalAmount float64 `json:"total_amount" db:"total_amount"`
	Status      string  `json:"status"`
	CreatedAt   string  `json:"created_at" db:"created_date"`
	UserId      int     `json:"user_id" db:"user_id"`
	Address     string  `json:"address"`
	ProductId   int     `json:"product_id" db:"product_id"`
	Name        string  `json:"name"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

type Status struct {
	StatusId int `json:"status_id"`
}
