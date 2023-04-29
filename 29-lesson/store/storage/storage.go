package storage

type StoreI interface {
	Store() StoreRepoI
}

type StoreRepoI interface {
	GetProduct(query []string) (err error)
}
