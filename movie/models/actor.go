package models

import "time"

type Actors struct {
	ID        string     `json:"id"`
	FirstName string     `json:"firstname"`
	LastName  string     `json:"lastname"`
	Age       int        `json:"age"`
	Country   string     `json:"country"`
	CreatedAT *time.Time `json:"created_at"`
	UpdatedAT *time.Time `json:"updated_at"`
}

type CreateActor struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
	Country   string `json:"country"`
}

type UpdateActor struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
	Country   string `json:"country"`
}
