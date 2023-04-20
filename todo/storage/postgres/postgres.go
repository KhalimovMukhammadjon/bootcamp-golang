package postgres

import (
	"Muhammadjon/bootcamp/todo/storage"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	db   *sqlx.DB
	todo storage.TodoRepoI
}

func NewPostgres(psqlConnString string) storage.StorageI {
	db, err := sqlx.Connect("postgres", psqlConnString)
	if err != nil {
		log.Fatal(err)
	}

	return &Store{
		db: db,
	}
}

func (s *Store) Todo() storage.TodoRepoI {
	if s.todo == nil {
		s.todo = &todoRepo{db: s.db}
	}
	return s.todo
}
