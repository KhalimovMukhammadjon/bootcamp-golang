package postgres

import (
	"Muhammadjon/bootcamp/movie/models"
	"Muhammadjon/bootcamp/movie/storage"
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type movieRepo struct {
	db *sqlx.DB
}

func NewMovieRepo(db *sqlx.DB) storage.MovieRepoI {
	return movieRepo{
		db: db,
	}
}

// Title       string  `json:"title"`
//     Description string  `json:"description"`
//     Year        int     `json:"year"`
//     ActorId     int     `json:"actor_id"`
//     Rating

func (r movieRepo) Create(entity models.CreateMovie) (err error) {
	insert := `INSERT INTO movie (
		title,
		description,
		year,
		actor_id,
		rating
	) VALUES(
		$1,
		$2,
		$3,
		$4,
		$5
	)`

	_, err = r.db.Exec(insert,
		entity.Title,
		entity.Description,
		entity.Year,
		entity.ActorId,
		entity.Rating,
	)

	return err
}

func (r movieRepo) GetMovieList(query models.Query) (resp []models.Movie, err error) {
	var rows *sql.Rows

	if len(query.Search) > 0 {
		rows, err = r.db.Query(
			`
			SELECT movie.id, movie.title, movie.description, movie.year, movie.rating,
			actor.firstname, actor.lastname, actor.age
			FROM movie
			JOIN actor ON movie.id = actor.id
			WHERE movie ILIKE '%' || $3 || '%'
			OFFSET $1 LIMIT $2
			`,
			query.Offset,
			query.Limit,
			query.Search,
		)
	} else {
		rows, err = r.db.Query(
			`SELECT movie.id, movie.title, movie.description, movie.year, movie.rating,
			actor.firstname, actor.lastname, actor.age
			FROM movie
			JOIN actor ON movie.id = actor.id
			WHERE movie ILIKE '%' || $3 || '%'
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
		var a models.Movie
		err = rows.Scan(
			&a.ID, &a.Title, &a.Description, &a.Year, &a.Rating,
			&a.Actor.ID, &a.Actor.FirstName, &a.Actor.LastName, &a.Actor.Age,
		)
		resp = append(resp, a)
		if err != nil {
			return resp, err
		}
	}

	return resp, err
}

func (r movieRepo) GetById(Id int) (resp models.Movie, err error) {
	return resp, err
}

func (r movieRepo) Update(entity models.UpdateMovie) (effectedRowsNum int64, err error) {
	return effectedRowsNum, err
}

func (r movieRepo) Delete(Id int) (effectedRowsNum int64, err error) {
	return effectedRowsNum, err
}
