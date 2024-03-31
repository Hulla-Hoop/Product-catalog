package db

import "testinhousead/internal/model"

// интерфейс каталога товаров на уровне бд
type DB interface {
	CreateCategory(reqId string, name string) (*model.Category, error)
	DeleteCategory(reqId string, id int) (*model.Category, error)
	UpdateCategory(reqId string, id int, name string) (*model.Category, error)
	AllCategories(reqID string) ([]model.Category, error)
	GoodsOnCateory(reqID string, category string) ([]model.Product, error)
	CreateGoods(reqId string, name string) (*model.Goods, error)
	DeleteGoods(reqId string, id int) (*model.Goods, error)
	UpdateGoods(reqId string, id int, name string) (*model.Goods, error)
	CreateRelation(reqId string, goods_id int, category_id int) (*model.Product, error)
	Close() error
}

// интерфейс авторизации
type DBAut interface {
	DeleteSess(reqId string, token string) error
	DeleteOld()
	CreateSess(reqId string, session *model.Session) error
	ChekSess(reqId string, token string) (*model.Session, error)
}
