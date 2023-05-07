package storage

import "Muhammadjon/bootcamp/movie/models"

type StorageI interface {
	Movie() MovieRepoI
	Actor() ActorRepoI
}

type ActorRepoI interface {
	Create(entity models.CreateActor) (err error)
	GetList(query models.Query) (resp []models.Actors, err error)
	GetById(Id string) (resp models.Actors, err error)
	Update(entity models.UpdateActor) (effectedRowsNum int64, err error)
	Delete(Id string) (effectedRowsNum int64, err error)
}

type MovieRepoI interface {
	Create(entity models.CreateMovie) (err error)
	GetMovieList(query models.Query) (resp []models.Movie, err error)
	GetById(Id int) (resp models.Movie, err error)
	Update(entity models.UpdateMovie) (effectedRowsNum int64, err error)
	Delete(Id int) (effectedRowsNum int64, err error)
}
