package postgres

import (
	"Muhammadjon/bootcamp/30-lesson/movie/config"
	"Muhammadjon/bootcamp/30-lesson/movie/storage"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	db    *pgxpool.Pool
	movie storage.MovieRepoI

	// actor storage.ActorRepoI
	// genre storage.GenresRepoI
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	//config.MaxConns = cfg.PostgresMaxConnections  //use

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx) // check if connection is alive
	if err != nil {
		return nil, err
	}

	return &Store{
		db: pool,
	}, err
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (l *Store) Log(ctx context.Context, msg string, data map[string]interface{}) {
	args := make([]interface{}, 0, len(data)+2) // making space for arguments + msg
	args = append(args, msg)
	for k, v := range data {
		args = append(args, fmt.Sprintf("%s=%v", k, v))
	}
	log.Println(args...)
}

func (s *Store) Movie() storage.MovieRepoI {
	if s.movie == nil {
		s.movie = NewMovieRepo(s.db)
	}
	return s.movie
}

// func (s *Store) Actor() storage.ActorRepoI {
// 	if s.actor == nil {
// 		s.actor = NewActorRepo(s.db)
// 	}

// 	return s.actor
// }

// func (s *Store) Genre() storage.GenresRepoI {
// 	if s.genre == nil {
// 		s.genre = NewGenreRepo(s.db)
// 	}

// 	return s.genre
