package storage

import (
	"Muhammadjon/bootcamp/30-lesson/movie/models"
	"context"
)

type StorageI interface {
	CloseDB()
	Movie() MovieRepoI
	// Actor() ActorRepoI
	// Genre() GenresRepoI
}

type MovieRepoI interface {
	Create(ctx context.Context, req models.CreateMovie) (resp string, error error)
	Update(id string, req models.UpdateMovie) error
	Delete(id string) error
	Get(id string) (resp models.Movie, error error)
	GetAll() (resp []models.Movie, error error)
}

// type ActorRepoI interface {
// 	Create(req models.CreateActor) (string, error)
// 	Update(id string, req models.UpdateActor) error
// 	Delete(id string) error
// 	Get(id string) (models.Actors, error)
// 	GetAll() ([]models.Actors, error)
// }

// type GenresRepoI interface {
// 	Create(req models.CreateGenre) (string, error)
// 	Update(id string, req models.UpdateGenre) error
// 	Delete(id string) error
// 	Get(id string) (models.Genres, error)
// 	GetAll() ([]models.Genres, error)
// }
