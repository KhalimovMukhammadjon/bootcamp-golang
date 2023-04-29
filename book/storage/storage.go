package storage

import "Muhammadjon/bootcamp/book/models"

type StorageI interface {
	Book() BookRepoI
}

type BookRepoI interface {
	Create(entity models.CreateBook) (err error)
	GetList() (resp []models.Book, err error)
	GetListById(ID string) (resp models.Book, err error)
	UpdateList(entity models.UpdateBook) (effectedRowsNum int64, err error)
	DeleteList(ID string) (effectedRowsNum int64, err error)
}

// type MovieRepoI interface {
// 	Create(ctx context.Context, req models.CreateMovie) (resp string, error error)
// 	Update(id string, req models.UpdateMovie) error
// 	Delete(id string) error
// 	Get(id string) (resp models.Movie, error error)
// 	GetAll() (resp []models.Movie, error error)
// }
