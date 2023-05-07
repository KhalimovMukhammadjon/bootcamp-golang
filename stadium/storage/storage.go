package storage

import "Muhammadjon/bootcamp/stadium/models"

type StorageI interface {
	Stadium() StadiumRepoI
}

type StadiumRepoI interface {
	Create(entity models.CreateStadium) (err error)
	GetStadiumList(query models.Query) (resp []models.Stadium, err error)
	GetById(Id int) (resp models.Stadium, err error)
	Update(entity models.UpdateStadium) (effectedRowsNum int64, err error)
	Delete(Id int) (effectedRowsNum int64, err error)
}
