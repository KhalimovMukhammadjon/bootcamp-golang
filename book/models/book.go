package models

import "time"

type Book struct {
	Id          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Created_at  *time.Time `json:"created_at"`
	Updated_at  *time.Time `json:"updated_at"`
}

type CreateBook struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateBook struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
