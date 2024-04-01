# API Каталог товаров

## Описание API
 
Представляет собой простой каталог товаров с возможность получения всех категорий и получения товаров по отдельным категориям.
Также аутифицированные пользователи могут добавлять,удалять,обновлять товары и категории.


## Архитектура и Паттерны

В проекте была применена `чистая архитектура`.

![Чистая архитектура](/images/clean.png)

### Сервис Каталога 
Отвечает за все взаимодействия с товарами и категориями 

собирается в pkg/app

```
...
type first struct {
	hand handlers.Catalog
}

func NewApp() (*app, error) {

	//обьявляем первый сервис и подключаем зависимости

	first := first{}
	l := logger.New()
	db, err := psql.InitDb(l)
	if err != nil {
		return nil, err
	}
	s := catalog.NewCatalog(l, db)
	first.hand = hand.NewCatalog(l, s) 
 ...
}

...
```

#### Handlers Слой 
Взаимодействует с клиентом(http) получает данные и передает их на сервисный слой по интерфейсу 

Интерфейс Handlers слоя
```
type Catalog interface {
	CreateCategory(w http.ResponseWriter, r *http.Request)
	DeleteCategory(w http.ResponseWriter, r *http.Request)
	UpdateCategory(w http.ResponseWriter, r *http.Request)
	AllCategories(w http.ResponseWriter, r *http.Request)
	GoodsOnCateory(w http.ResponseWriter, r *http.Request)
	CreateGoods(w http.ResponseWriter, r *http.Request)
	DeleteGoods(w http.ResponseWriter, r *http.Request)
	UpdateGoods(w http.ResponseWriter, r *http.Request)
}
```
Имплементация интерфейса лежит в *handlers/catalog*

#### Service Слой 
В этом слое содержится вся бизнес логика.

Получает данные от Handlers Слоя преобразует их в нужный формат и передает в слой бд при необходимости все последующие проверки данных будут реализованны в этом слое.

Интерфейс Service слоя 
``` 
type Cataloger interface {
	CreateCategory(reqId string, name string) ([]byte, error)
	DeleteCategory(reqId string, id string) ([]byte, error)
	UpdateCategory(reqId string, id string, name string) ([]byte, error)
	AllCategories(reqId string) ([]byte, error)
	GoodsOnCateory(reqId string, category string) ([]byte, error)
	CreateGoods(reqId string, name string, category string) ([]byte, error)
	DeleteGoods(reqId string, id string) ([]byte, error)
	UpdateGoods(reqId string, id string, name string) ([]byte, error)
	Close() error
}
```
Имплементация интерфейса лежит в *service/catalog*

#### DB Слой 
Отвечает за взаимодействие с базой данных 

Интерфейс DB слоя 
```
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
```

Имплементация интерфейса лежит в *DB/psql* 

