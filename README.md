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
}```
