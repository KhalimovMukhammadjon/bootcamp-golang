package postgres

import (
	"Muhammadjon/bootcamp/movie/models"
	"Muhammadjon/bootcamp/movie/storage"
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type actorRepo struct {
	db *sqlx.DB
}

func NewActorRepo(db *sqlx.DB) storage.ActorRepoI {
	return actorRepo{
		db: db,
	}
}

func (r actorRepo) Create(entity models.CreateActor) (err error) {
	insert := `INSERT INTO actor (
		firstname,
		lastname,
		age,
		country
	) VALUES(
		$1,
		$2,
		$3,
		$4
	)`

	_, err = r.db.Exec(insert,
		entity.FirstName,
		entity.LastName,
		entity.Age,
		entity.Country,
	)

	return err
}

func (r actorRepo) GetList(query models.Query) (resp []models.Actors, err error) {
	var rows *sql.Rows

	if len(query.Search) > 0 {
		rows, err = r.db.Query(
			`
			SELECT id, firstname, lastname, age, country, created_at, updated_at
			FROM actor
			WHERE firstname ILIKE '%' || $3 || '%'
			OFFSET $1 LIMIT $2
		`,
			query.Offset,
			query.Limit,
			query.Search,
		)
	} else {
		rows, err = r.db.Query(
			`
			SELECT id, firstname, lastname, age, country, created_at, updated_at
			FROM actor
			OFFSET $1 LIMIT $2
		`,
			query.Offset,
			query.Limit,
		)
	}

	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var r models.Actors
		err = rows.Scan(
			&r.ID, &r.FirstName, &r.LastName, &r.Age, &r.Country, &r.CreatedAT, &r.UpdatedAT,
		)
		resp = append(resp, r)
		if err != nil {
			return resp, err
		}
	}

	return resp, err

}

func (r actorRepo) GetById(Id string) (resp models.Actors, err error) {
	query, err := r.db.Prepare(`
		SELECT id, firstname, lastname, age, country, created_at, updated_at
		FROM actor
		WHERE id = $1
	`,
	)
	if err != nil {
		return resp, err
	}
	defer query.Close()

	err = query.QueryRow(Id).Scan(
		&resp.ID, &resp.FirstName, &resp.LastName, &resp.Age, &resp.Country, &resp.CreatedAT, &resp.UpdatedAT,
	)

	return resp, err
}

func (r actorRepo) Update(entity models.UpdateActor) (effectedRowsNum int64, err error) {
	query := `UPDATE actor SET 
		firstname = $2,
		lastname = $3,
		age = $4,
		country = $5,
		updated_at = now()
		WHERE id = $1
	`
	_, err = r.db.Exec(query,
		entity.ID,
		entity.FirstName,
		entity.LastName,
		entity.Age,
		entity.Country,
	)
	if err != nil {
		return effectedRowsNum, err
	}

	return effectedRowsNum, nil
}

func (r actorRepo) Delete(Id string) (effectedRowsNum int64, err error) {
	query := `DELETE FROM actor WHERE id = $1`
	if err != nil {
		return effectedRowsNum, err
	}

	result, err := r.db.Exec(query,
		Id,
	)
	if err != nil {
		return effectedRowsNum, err
	}

	effectedRowsNum, err = result.RowsAffected()
	if err != nil {
		return effectedRowsNum, err
	}

	return effectedRowsNum, nil
}
