package models

import "time"

type Movie struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Year        int        `json:"year"`
	Genre       string     `json:"genre"`
	Description string     `json:"description"`
	ImageURL    string     `json:"image_url"`
	Rating      float32    `json:"rating"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type CreateMovie struct {
	Title       string  `json:"title"`
	Year        int     `json:"year"`
	Genre       string  `json:"genre"`
	Description string  `json:"description"`
	ImageURL    string  `json:"image_url"`
	Rating      float32 `json:"rating"`
}

type UpdateMovie struct {
	Title       string  `json:"title"`
	Year        int     `json:"year"`
	Genre       string  `json:"genre"`
	Description string  `json:"description"`
	ImageURL    string  `json:"image_url"`
	Rating      float32 `json:"rating"`
}
