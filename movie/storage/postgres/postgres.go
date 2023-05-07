package postgres

import (
	"Muhammadjon/bootcamp/movie/storage"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	db    *sqlx.DB
	actor storage.ActorRepoI
	movie storage.MovieRepoI
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

func (s *Store) Actor() storage.ActorRepoI {
	if s.actor == nil {
		s.actor = &actorRepo{db: s.db}
	}
	return s.actor
}

func (s *Store) Movie() storage.MovieRepoI {
	if s.movie == nil {
		s.movie = &movieRepo{db: s.db}
	}
	return s.movie
}
