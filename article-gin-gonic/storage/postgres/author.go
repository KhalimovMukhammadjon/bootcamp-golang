package postgres

import (
	"bootcamp/article/models"
	"bootcamp/article/storage"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type authorRepo struct {
	db *sqlx.DB
}

func NewAuthorRepo(db *sqlx.DB) storage.AuthorRepoI {
	return authorRepo{
		db: db,
	}
}

func (r authorRepo) Create(entity models.PersonCreateModel) (err error) {
	insertQuery := `INSERT INTO article (
		title,
		body,
	) VALUES (
		$1, -- 
		$2, -- 
	)`

	_, err = r.db.Exec(insertQuery,
		entity.Firstname,
		entity.Lastname,
	)

	return err
}

func (r authorRepo) GetList(query models.Query) (resp []models.Person, err error) {
	var rows *sql.Rows
	//t.db.QueryRow()

	if len(query.Search) > 0 {
		rows, err = r.db.Query(
			`
			SELECT id, firstname, lastname, age, gender, created_at, updated_at from author
			WHERE title ILIKE '%' || $3 || '%'
			OFFSET $1 LIMIT $2
			`,
			query.Offset,
			query.Limit,
			query.Search,
		)
	} else {
		rows, err = r.db.Query(
			`SELECT  id, firstname, lastname, age, gender, created_at, updated_at from author
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
		var r models.Person
		err = rows.Scan(
			&r.ID, &r.Firstname, &r.Lastname, &r.Age, &r.Gender, &r.CreatedAt, &r.UpdatedAt,
		)
		resp = append(resp, r)
		if err != nil {
			return resp, err
		}
	}
	return resp, err
}

func (r authorRepo) GetByID(ID int) (resp models.Person, err error) {
	var rows *sql.Rows

	query, err := r.db.Prepare(`
		SELECT id, firstname, lastname, age, gender, created_at, updated_at FROM author
		WHERE id = $1 
	`)

	if err != nil {
		return resp, err
	}
	defer query.Close()

	err = rows.Scan(&resp.ID, &resp.Firstname, &resp.Lastname, &resp.Age, &resp.Gender, &resp.CreatedAt, &resp.UpdatedAt)
	return resp, err
}

func (r authorRepo) Update(entity models.PersonUpdateModel) (effectedRowsNum int, err error) {
	query := `UPDATE author SET
		firstname = $2
		lastname = $3
		age = $4
		updated_at = now()
			WHERE id = $1
	`
	_, err = r.db.Exec(query,
		entity.ID,
		entity.Firstname,
		entity.Lastname,
		entity.Age,
	)
	if err != nil {
		return effectedRowsNum, err
	}
	return effectedRowsNum, nil
}

func (r authorRepo) Delete(ID string) (effectedRowsNum int64, err error) {
	delete := `DELETE FROM author where id = $1`
	if err != nil {
		return effectedRowsNum, err
	}

	result, err := r.db.Exec(delete, ID)
	if err != nil {
		return effectedRowsNum, err
	}

	effectedRowsNum, err = result.RowsAffected()
	if err != nil {
		return effectedRowsNum, err
	}
	return effectedRowsNum, nil
}

func (r authorRepo) GetMostPublisher() (resp []models.MostPublisher, err error) {
	var rows *sql.Rows

	rows, err = r.db.Query(`
		SELECT author.firstname, COUNT(article.author_id)
		FROM author
		JOIN article ON article.author_id = author.id
		GROUP BY author.id ORDER BY count DESC limit 3
		`)

	if err != nil {
		return resp, err
	}

	defer rows.Close()

	for rows.Next() {
		var r models.MostPublisher
		err = rows.Scan(
			&r.Firstname, &r.Count,
		)
		resp = append(resp, r)
		if err != nil {
			return resp, err
		}
	}
	return resp, err
}

func (r authorRepo) GetMinPublisher() (resp []models.MostPublisher, err error) {
	var rows *sql.Rows

	rows, err = r.db.Query(`
		SELECT author.firstname, COUNT(article.author_id)
		FROM author
		JOIN article ON article.author_id = author.id
		GROUP BY author.id ORDER BY count ASC limit 2
	`)

	if err != nil {
		return resp, err
	}

	defer rows.Close()

	for rows.Next() {
		var r models.MostPublisher
		err = rows.Scan(
			&r.Firstname,
			&r.Count,
		)
		resp = append(resp, r)
		if err != nil {
			return resp, err
		}
	}

	return resp, err
}
