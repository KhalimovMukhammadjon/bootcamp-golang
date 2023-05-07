package models

type User struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Phone_number int    `json:"phone_number"`
}

type CreateUser struct {
	Name         string `json:"name"`
	Phone_number int    `json:"phone_number"`
}

type UpdateAUser struct {
	Name         string `json:"name"`
	Phone_number int    `json:"phone_number"`
}
