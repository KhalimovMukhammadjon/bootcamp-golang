package postgres

import (
	"Muhammadjon/bootcamp/todo/models"
	"Muhammadjon/bootcamp/todo/storage"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type todoRepo struct {
	db   *sqlx.DB
	todo storage.TodoRepoI
}

func (t todoRepo) Create(entity models.CreateTodo) (err error) {
	insertQuery := `INSERT INTO todo(
		title,
		description
	) VALUES(
		$1,
		$2
	)`

	_, err = t.db.Exec(insertQuery,
		entity.Title,
		entity.Description,
	)
	return err
}

// GET list
func (t todoRepo) GetList(query models.Query) (resp []models.Todo, err error) {
	var rows *sql.Rows

	if len(query.Search) > 0 {
		rows, err = t.db.Query(
			`
			SELECT id, title, description, completed, created_at, updated_at from todo
			WHERE title ILIKE '%' || $3 || '%'
			OFFSET $1 LIMIT $2
			`,
			query.Offset,
			query.Limit,
			query.Search,
		)
	} else {
		rows, err = t.db.Query(
			`SELECT id, title, description, completed, created_at, updated_at from todo
			OFFSET $1 LIMIT $2`,
			query.Offset,
			query.Limit,
		)
	}

	if err != nil {
		return resp, err
	}

	defer rows.Close()

	for rows.Next() {
		var r models.Todo
		err = rows.Scan(
			&r.Id, &r.Title, &r.Description, &r.Completed, &r.Created_at, &r.Updated_at,
		)
		resp = append(resp, r)
		if err != nil {
			return resp, err
		}
	}
	return resp, err
}

func (t todoRepo) GetListById(ID string) (resp models.Todo, err error) {
	get, err := t.db.Prepare(`SELECT id, title, description, completed, created_at, updated_at 
	from todo
	WHERE id = $1`)
	if err != nil {
		return resp, err
	}
	defer get.Close()
	err = get.QueryRow(ID).Scan(&resp.Id, &resp.Title, &resp.Description, &resp.Completed, &resp.Created_at, &resp.Updated_at)
	return resp, err
}

func (t todoRepo) DeleteList(ID string) (effectedRowsNum int64, err error) {

	delete := `DELETE FROM todo WHERE id=$1`
	if err != nil {
		return effectedRowsNum, err
	}

	result, err := t.db.Exec(delete, ID)
	if err != nil {
		return effectedRowsNum, err
	}

	// Get the number of affected rows
	effectedRowsNum, err = result.RowsAffected()
	if err != nil {
		return effectedRowsNum, err
	}
	return effectedRowsNum, nil
}

func (t todoRepo) UpdateList(entity models.UpdateTodo) (effectedRowsNum int, err error) {
	return
}

// func (t todoRepo) DeleteList(ID string) (effectedRowsNum int64, err error) {

// 	delete, err := t.db.Prepare(`DELETE FROM todo WHERE id = $1`)
// 	if err != nil {
// 		return effectedRowsNum, err
// 	}
// 	defer delete.Close()

// 	result, err := t.db.Exec(ID)
// 	if err != nil {
// 		return effectedRowsNum, err
// 	}

// 	// Get the number of affected rows
// 	effectedRowsNum, err = result.RowsAffected()
// 	if err != nil {
// 		return effectedRowsNum, err
// 	}

// 	return effectedRowsNum, err
// }
