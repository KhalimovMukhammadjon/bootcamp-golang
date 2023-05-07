package postgres

import (
	//"Muhammadjon/bootcamp/movie/storage"

	"Muhammadjon/bootcamp/stadium/models"
	"Muhammadjon/bootcamp/stadium/storage"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/jmoiron/sqlx"
	// "github.com/jackc/pgtype"
	_ "github.com/lib/pq"
)

type stadiumRepo struct {
	db *sqlx.DB
}

func NewStadiumRepo(db *sqlx.DB) storage.StadiumRepoI {
	return stadiumRepo{
		db: db,
	}
}

func (r stadiumRepo) Create(entity models.CreateStadium) (err error) {

	insert := `INSERT INTO stadium(
		stadium_name,
		available_time_start,
		available_time_end,
		price_p_hour,
		rating,
		images,
		location
	  ) VALUES(
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7
	  )`

	loc, err := json.Marshal(entity.Location)
	if err != nil {
		log.Fatal(err)
	}

	_, err = r.db.Exec(insert,
		entity.StadiumName,
		entity.AvailableTimeStart,
		entity.AvailableTimeEnd,
		entity.PricePerHour,
		entity.Rating,
		entity.Images,
		loc,
	)
	return err 
}

func (r stadiumRepo) GetStadiumList(query models.Query) (resp []models.Stadium, err error) {
	var rows *sql.Rows

	if len(query.Search) > 0 {
		rows, err = r.db.Query(
			`
			SELECT id, stadium_name, available_time_start, available_time_end, price_p_hour,rating,images
			,location,created_at, updated_at
			FROM stadium
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
			SELECT id, stadium_name, available_time_start, available_time_end, price_p_hour,rating,images
			,location,created_at, updated_at
			FROM stadium
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
		var r models.Stadium

		// jsonString, err := json.Marshal(r.Location)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		err = rows.Scan(
			&r.ID, &r.StadiumName, &r.AvailableTimeStart, &r.AvailableTimeEnd,
			&r.PricePerHour, &r.Rating, &r.Images, &r.Location, &r.CreatedAt, &r.UpdatedAt,
		)
		resp = append(resp, r)
		if err != nil {
			return resp, err
		}
	}
	return resp, err
}
func (r stadiumRepo) GetById(Id int) (resp models.Stadium, err error) {
	// query, err := r.db.Prepare(`
	// 	SELECT id, firstname, lastname, age, country, created_at, updated_at
	// 	FROM actor
	// 	WHERE id = $1
	// `,
	// )
	// if err != nil {
	// 	return resp, err
	// }
	// defer query.Close()

	// err = query.QueryRow(Id).Scan(
	// 	&resp.ID, &resp.FirstName, &resp.LastName, &resp.Age, &resp.Country, &resp.CreatedAT, &resp.UpdatedAT,
	// )

	return resp, err
}

func (r stadiumRepo) Update(entity models.UpdateStadium) (effectedRowsNum int64, err error) {
	return effectedRowsNum, err
}
func (r stadiumRepo) Delete(Id int) (effectedRowsNum int64, err error) {
	return effectedRowsNum, err
}
