package storage

type Product struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Price       float32             `json:"price"`
	Attributes  map[string][]string `json:"attributes"`
}

type Store struct {
	Products []Product
}

type ProductV struct {
	Products Product
}
