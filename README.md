# API Каталог товаров

## Описание API
 
Представляет собой простой каталог товаров с возможность получения всех категорий и получения товаров по отдельным категориям.
Также аутифицированные пользователи могут добавлять,удалять,обновлять товары и категории.


## Архитектура

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
Имплементация интерфейса лежит в *internal/handlers/catalog*

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
Имплементация интерфейса лежит в *internal/service/catalog*

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

Имплементация интерфейса лежит в *internal/DB/psql* 

#### Структура БД 

- В качестве БД выбрана PostgreSQL 
- БД соответствует 4 нормальной форме
- БД создается по средством миграций *migrations/*
- Также миграции добавляют несколько записей при вызове для удобства тестирования

|Name|Category|
|:-:|:-:|
|Spagetti|For boil|
|Spagetti|Flour products|
|Apple|Fruits|
|Apple|Fresh|
|Apple Pie|Bakery products|
|Apple Pie|Fruits|

`Схема БД`

![Схема Базы данных](/images/CatalogDB.png)

### Сервис Аутификации

Отвечает за аутификацию пользователей и хранение сессий пользователей

Пользователи доступные для аутификации лежат в *internal/model/modelAut*
```
 var Users = map[string]string{
	"3825c945-8843-4b7d-995e-30b16c173c65": "user1",
	"019ed7ca-8286-40b8-ac80-1950c92dccfd": "user2",
}
```

собирается в pkg/app

```
...
type second struct {
	hand handlers.Aut
}

func NewApp() (*app, error) {
...
	//обьявляем второй сервис и подключаем зависимости
	second := second{}
	mongo := mongo.New(l)
	se := autification.NewAut(l, mongo)
	second.hand = aut.NewAut(l, se)
 ...
}

...
```

#### Handlers Слой 
Взаимодействует с клиентом(http) получает данные и передает их на сервисный слой по интерфейсу 

Интерфейс Handlers слоя
```
type Aut interface {
	SignIn(w http.ResponseWriter, r *http.Request)
	Refresh(w http.ResponseWriter, r *http.Request)
}
```
Имплементация интерфейса лежит в *internal/handlers/autification*

#### Service Слой 
В этом слое содержится бизнес-логика(проверка наличия пользователей расшифровка токенов и тд).

Получает данные от Handlers Слоя преобразует их в нужный формат проверяет на корректность и передает в слой бд.

Интерфейс Service слоя 
``` 
type Autificationer interface {
	GetTokens(reqId string, guid string) (*http.Cookie, *http.Cookie, error)
	RefreshToken(reqID string, token string) (bool, string)
	ChekSess(reqId string, token string) (*model.Session, error)
}
```
Имплементация интерфейса лежит в *internal/service/autification*

#### DB Слой 
Отвечает за взаимодействие с базой данных (mongo)

Интерфейс DB слоя 
```
type DBAut interface {
	DeleteSess(reqId string, token string) error
	DeleteOld()
	CreateSess(reqId string, session *model.Session) error
	ChekSess(reqId string, token string) (*model.Session, error)
}
```

Имплементация интерфейса лежит в *internal/DB/mongo* 

#### Структура БД 

- В качестве БД для хранения сессий используется mongo
- Документы в mongo соответствуют структуре Session

```
type Session struct {
	BcryptTocken      string `bson:"bcryptTocken"`
	TimeCreatedTocken string `bson:"timeCreatedTocken"`
	Guid              string `bson:"guid"`
	ExpireTime        int64  `bson:"expiretime"`
}

```

##  