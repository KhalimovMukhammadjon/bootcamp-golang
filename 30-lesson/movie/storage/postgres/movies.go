package postgres

import (
	"Muhammadjon/bootcamp/30-lesson/movie/models"
	"Muhammadjon/bootcamp/30-lesson/movie/storage"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type movieRepo struct {
	db *pgxpool.Pool
}

func NewMovieRepo(db *pgxpool.Pool) storage.MovieRepoI {
	return &movieRepo{
		db: db,
	}
}

func (m *movieRepo) Create(ctx context.Context, req models.CreateMovie) (resp string, error error) {
	query := `insert into movies
			  (title, year, genre, description, image_url, rating)
			  values ($1, $2, $3, $4, $5, $6)`

	_, err := m.db.Exec(ctx, query,
		req.Title,
		req.Year,
		req.Genre,
		req.Description,
		req.ImageURL,
		req.Rating,
	)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (m *movieRepo) Update(id string, req models.UpdateMovie) error {
	return nil
}

func (m *movieRepo) Delete(id string) error {
	return nil
}

func (m *movieRepo) Get(id string) (resp models.Movie, error error) {
	return models.Movie{}, nil
}

func (m *movieRepo) GetAll() (resp []models.Movie, error error) {

	return []models.Movie{}, nil
}
