package storage

import "Muhammadjon/bootcamp/todo/models"

type StorageI interface {
	Todo() TodoRepoI
}

// Create, Update, Delete, Get
type TodoRepoI interface {
	Create(entity models.CreateTodo) (err error)
}
