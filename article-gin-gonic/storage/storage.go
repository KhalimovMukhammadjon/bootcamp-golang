package storage

import "bootcamp/article/models"

type StorageI interface {
	Article() ArticleRepoI
	Author() AuthorRepoI
}

type ArticleRepoI interface {
	Create(entity models.ArticleCreateModel) (err error)
	GetList(query models.Query) (resp []models.ArticleListItem, err error)
	GetByID(ID int) (resp models.Article, err error)
	Update(entity models.ArticleUpdateModel) (effectedRowsNum int, err error)
	Delete(ID string) (effectedRowsNum int64, err error)
}

type AuthorRepoI interface {
	Create(entity models.PersonCreateModel) (err error)
	GetList(query models.Query) (resp []models.Person, err error)
	GetByID(ID int) (resp models.Person, err error)
	Update(entity models.PersonUpdateModel) (effectedRowsNum int, err error)
	Delete(ID string) (effectedRowsNum int64, err error)
	GetMostPublisher() (resp []models.MostPublisher, err error)
	GetMinPublisher() (resp []models.MostPublisher, err error)
}
