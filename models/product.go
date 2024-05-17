package models

type Product struct {
	Id                  int      `json:"-" db:"id"`
	Name                string   `json:"name" binding:"required"`
	Image               string   `json:"image" binding:"required"`
	Price               float64  `json:"price" binding:"required"`
	Height              string   `json:"height" binding:"required"`
	Size                string   `json:"size" binding:"required"`
	Instruction         string   `json:"instruction"`
	Description         string   `json:"description" binding:"required"`
	RecommendedProducts []string `json:"recommended_products"`
}

type GetProducts struct {
	Id    int     `json:"id" db:"id"`
	Name  string  `json:"name"`
	Image string  `json:"image"`
	Price float64 `json:"price"`
}

type GetProductsFromCart struct {
	Id       int     `json:"id" db:"id"`
	Name     string  `json:"name"`
	Image    string  `json:"image"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type GetProduct struct {
	Id          int     `json:"id" db:"id"`
	Name        string  `json:"name"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Height      string  `json:"height"`
	Size        string  `json:"size"`
	Instruction string  `json:"instruction"`
	Description string  `json:"description"`
}

type UpdateProduct struct {
	Name                string   `json:"name"`
	Image               string   `json:"image"`
	Price               float64  `json:"price"`
	Height              string   `json:"height"`
	Size                string   `json:"size"`
	Instruction         string   `json:"instruction"`
	Description         string   `json:"description"`
	RecommendedProducts []string `json:"recommended_products"`
}

type Address struct {
	Address string `json:"address"`
}
