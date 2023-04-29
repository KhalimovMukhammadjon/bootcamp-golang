package postgres

import (
	"bootcamp/article/models"
	"bootcamp/article/storage"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type articleRepo struct {
	db *sqlx.DB
}

func NewArticleRepo(db *sqlx.DB) storage.ArticleRepoI {
	return articleRepo{
		db: db,
	}
}

func (r articleRepo) Create(entity models.ArticleCreateModel) (err error) {
	insertQuery := `INSERT INTO article (
		title,
		body,
		author_id
	) VALUES (
		$1, -- 
		$2, -- 
		$3 --
	)`

	_, err = r.db.Exec(insertQuery,
		entity.Title,
		entity.Body,
		entity.AuthorID,
	)

	return err
}

func (r articleRepo) GetList(query models.Query) (resp []models.ArticleListItem, err error) {
	var rows *sql.Rows
	if len(query.Search) > 0 {
		rows, err = r.db.Query(
			`SELECT
		ar.id, ar.title, ar.body, ar.created_at, ar.updated_at,
		au.id, au.firstname, au.lastname, au.created_at, au.updated_at
		FROM article AS ar
		JOIN author AS au ON ar.author_id = au.id
		WHERE title ILIKE '%' || $3 || '%' AND deleted_at IS NULL
		OFFSET $1 LIMIT $2
		`,
			query.Offset,
			query.Limit,
			query.Search,
		)

	} else {
		rows, err = r.db.Query(
			`SELECT
			ar.id, ar.title, ar.body, ar.created_at, ar.updated_at,
			au.id, au.firstname, au.lastname, au.created_at, au.updated_at
			FROM article AS ar JOIN author AS au ON ar.author_id = au.id WHERE deleted_at IS NULL
			OFFSET $1 LIMIT $2`, // offset 0 limit 10 -> 1-10,
			query.Offset,
			query.Limit,
		)
	}

	// query := "SELECT * FROM users WHERE deleted_at IS NULL LIMIT ? OFFSET ?"
	// rows, err := db.Query(query, limit, offset)
	// if err != nil {
	//     return nil, err
	// }

	if err != nil {
		return resp, err
	}

	defer rows.Close()
	for rows.Next() {
		var a models.ArticleListItem
		err = rows.Scan(
			&a.ID, &a.Title, &a.Body, &a.CreatedAt, &a.UpdatedAt,
			&a.Author.ID, &a.Author.Firstname, &a.Author.Lastname, &a.Author.CreatedAt, &a.Author.UpdatedAt,
		)
		resp = append(resp, a)
		if err != nil {
			return resp, err
		}
	}

	return resp, err
}

func (r articleRepo) GetByID(ID int) (resp models.Article, err error) {

	return
}

func (r articleRepo) Update(entity models.ArticleUpdateModel) (effectedRowsNum int, err error) {
	query := `UPDATE todo SET
	title = $2,
	body = $3,
	WHERE id = $1`

	_, err = r.db.Exec(query,
		entity.ID,
		entity.Title,
		entity.Body,
	)

	if err != nil {
		return effectedRowsNum, err
	}
	return effectedRowsNum, nil
}

func (r articleRepo) Delete(ID string) (effectedRowsNum int64, err error) {
	delete := `UPDATE article SET deleted_at = NOW() WHERE id = $1`
	if err != nil {
		return effectedRowsNum, err
	}

	query, err := r.db.Exec(delete,
		ID,
	)

	if err != nil {
		return effectedRowsNum, err
	}

	effectedRowsNum, err = query.RowsAffected()
	if err != nil {
		return effectedRowsNum, err
	}
	return effectedRowsNum, nil
}
