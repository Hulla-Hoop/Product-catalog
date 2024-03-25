package service

type MarketService interface {
	CreateCategory(reqId string, name string) ([]byte, error)
	DeleteCategory(reqId string, id string) ([]byte, error)
	UpdateCategory(reqId string, id string, name string) ([]byte, error)
	AllCategories(reqId string) ([]byte, error)
	GoodsOnCateory(reqId string, category string) ([]byte, error)
	CreateGoods(reqId string, name string) ([]byte, error)
	DeleteGoods(reqId string, id string) ([]byte, error)
	UpdateGoods(reqId string, id string, name string) ([]byte, error)
	Close() error
}
