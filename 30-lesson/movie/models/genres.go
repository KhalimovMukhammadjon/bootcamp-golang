package models

import "time"

type Genres struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	CreatedAT *time.Time `json:"created_at"`
	UpdatedAT *time.Time `json:"updated_at"`
}

type CreateGenre struct {
	Name string `json:"name"`
}

type UpdateGenre struct {
	Name string `json:"name"`
}

type MovieGenres struct {
	ID        int64      `json:"id"`
	MovieID   int64      `json:"movie_id"`
	GenreID   int64      `json:"genre_id"`
	CreatedAT *time.Time `json:"created_at"`
	UpdatedAT *time.Time `json:"updated_at"`
}
