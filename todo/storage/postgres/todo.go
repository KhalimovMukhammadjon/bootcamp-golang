package postgres

import (
	"Muhammadjon/bootcamp/todo/models"
	"Muhammadjon/bootcamp/todo/storage"

	"github.com/jmoiron/sqlx"
)

type todoRepo struct {
	db   *sqlx.DB
	todo storage.TodoRepoI
}

func (t todoRepo) Create(entity models.CreateTodo) (err error) {
	insertQuery := `INSERT INTO Todo(
		title,
		description
	) VALUES 
		$1,
		$2
	)`

	_, err = t.db.Exec(insertQuery,
		entity.Title,
		entity.Description,
	)
	return err
}
