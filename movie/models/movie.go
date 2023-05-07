package models

import "time"

type Movie struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Year        string     `json:"year"`
	Actor       Actors     `json:"actor"`
	Rating      float32    `json:"rating"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type CreateMovie struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Year        string  `json:"year"`
	ActorId     int     `json:"actor_id"`
	Rating      float32 `json:"rating"`
}

type UpdateMovie struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Year        string  `json:"year"`
	ImageURL    string  `json:"image_url"`
	Rating      float32 `json:"rating"`
}

/*type Movie struct {
	Movie_id Movie
	Actor_id Actor
 }
*/
