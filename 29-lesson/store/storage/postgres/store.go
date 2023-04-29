package postgres

import "Muhammadjon/bootcamp/29-lesson/store/storage"

type Store struct {
	//db   *sqlx.DB
	product storage.StoreRepoI
}

func (c Store) GetProduct(query []string) (err error) {
	return
}
