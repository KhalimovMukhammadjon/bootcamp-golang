package models

import "time"

// this is the comment
type Person struct {
	ID        string     `json:"id"`
	Firstname string     `json:"firstname"`
	Lastname  string     `json:"lastname"`
	Age       int32      `json:"age"`
	Gender    string     `json:"gender"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type PersonCreateModel struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type PersonUpdateModel struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int32  `json:"age"`
}

type MostPublisher struct {
	Firstname string `json:"firstname"`
	Count     int    `json:"count"`
}
