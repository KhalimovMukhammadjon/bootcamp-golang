package postgres

import (
	"Muhammadjon/bootcamp/book/models"
	//"Muhammadjon/bootcamp/book/storage"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type bookRepo struct {
	db *pgxpool.Pool
}

// func NewBookRepo(db *pgxpool.Pool) storage.BookRepoI {
// 	return *bookRepo{
// 		db: db,
// 	}
// }

func (t bookRepo) Create(ctx context.Context, entity models.CreateBook) (err error) {
	insertQuery := `INSERT INTO todo(
		title,
		description
	) VALUES(
		$1,
		$2
	)`

	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	_, err = t.db.Exec(ctx, insertQuery,
		uuid.String(),
		entity.Title,
		entity.Description,
	)

	if err != nil{
		
	}
	return err
}
