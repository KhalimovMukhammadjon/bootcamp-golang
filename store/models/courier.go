package models

type Courier struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Phone_number int    `json:"phone_number"`
}

type CreateCourier struct {
	Name         string `json:"name"`
	Phone_number int    `json:"phone_number"`
}

type UpdateCourier struct {
	Name         string `json:"name"`
	Phone_number int    `json:"phone_number"`
}
