package models

type Customer struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone int    `json:"phone"`
}

type CreateActor struct {
	Name  string `json:"name"`
	Phone int    `json:"phone"`
}

type UpdateActor struct {
	Name  string `json:"name"`
	Phone int    `json:"phone"`
}
