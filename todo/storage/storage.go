package storage

import "Muhammadjon/bootcamp/todo/models"

type StorageI interface {
	Todo() TodoRepoI
}

// Create, Update, Delete, Get
type TodoRepoI interface {
	Create(entity models.CreateTodo) (err error)
	GetList(query models.Query) (resp []models.Todo, err error)
	GetListById(ID string) (resp models.Todo, err error)
	UpdateList(entity models.UpdateTodo) (effectedRowsNum int, err error)
	DeleteList(ID string) (effectedRowsNum int64, err error)
}
