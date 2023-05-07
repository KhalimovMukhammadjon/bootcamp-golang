package models

type Product struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Price      float32 `json:"price"`
	CategoryID string  `json:"category_id"`
}

type CreateProduct struct {
	Name       string  `json:"name"`
	Price      float32 `json:"price"`
	CategoryID string  `json:"category_id"`
}

type UpdateProduct struct {
	Name       string  `json:"name"`
	Price      float32 `json:"price"`
	CategoryID string  `json:"category_id"`
}
