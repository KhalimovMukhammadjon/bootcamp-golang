package postgres

import (
	"Muhammadjon/bootcamp/stadium/storage"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	db      *sqlx.DB
	stadium storage.StadiumRepoI
}

func NewPostgres(psqlConnString string) storage.StorageI {
	db, err := sqlx.Connect("postgres", psqlConnString)
	if err != nil {
		log.Panic(err)
	}

	return &Store{
		db: db,
	}
}

func (s *Store) Stadium() storage.StadiumRepoI {
	if s.stadium == nil {
		s.stadium = &stadiumRepo{
			db: s.db,
		}
	}
	return s.stadium
}
