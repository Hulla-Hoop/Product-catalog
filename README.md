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
### Middleware
Обертки над ручками для добавления дополнительной логики в данном случае авторизация и реквест айди
#### Aut 
Проверяет есть ли у пользователя токен и соответствует ли он необходимым условиям 
``` 
func Aut(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID, ok := r.Context().Value("reqID").(string)
		if !ok {
			reqID = ""
		}

		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		tknStr := c.Value

		claims := &model.Claims{}

		cfg := config.TokenCFG()

		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (any, error) {
			return []byte(cfg.SecretKey), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "reqID", reqID)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
```

#### reqID 

Добавляет айди для запросов для удобного отслеживания логов в рамках одного запроса

``` 
func ReqID(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get("X-Request-ID")
		if reqID == "" {
			reqID = uuid.New().String()
		}
		ctx := context.WithValue(r.Context(), "reqID", reqID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

``` 
### Доп пакеты

#### Logger
В качестве логера используется библиотека Логрус объявляется один раз и пробрасывается везде.

#### Config
Выгружает необходимые конфигурации из .env 



## Endpoints

**Спецификация swagger** -- <https://github.com/Hulla-Hoop/Product-catalog/tree/main/docs>

### /allcategories/
возвращат все не удаленные категории
```
 /allcategories:
    post:
      tags:
        - allcategories
      summary: Возвращает все доступнные категории
      description: Возвращает все доступнные категории
      responses:
        '200':
          description: Successful operation
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/categories'          
        '500':
          description: Выводит ошибку
```

### /goodsoncategory
```  
/goodsoncategory:
    post:
      tags:
        - goodsoncategory
      summary: Возращает товары
      description: Возвращает товары по указанной категории
      parameters:
        - name: name
          in: query
          description: название категории
          required: true
          schema:
            type: string
            default: 
      responses:
        '200':
          description: Successful operation
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/products'
        '400':
          description: Пустой запрос
        '500':
          description: Выводит ошибку
```

### /category/create
```
/category/create:
    post:
      tags: 
      - category
      summary: Создает категорию
      description: Возращает созданую категорию
      parameters:
        - name: name
          in: query
          description: название категории
          required: true
          schema:
            type: string
            default: 
      responses:
        '200':
          description: Successful operation
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Category'
        '400':
          description: Пустой запрос
        '401':
          description: ""
        '500':
          description: Выводит ошибку
      security:
        - JwtCookieAuth: []
```
### /category/delete
``` 
/category/delete:
    post:
      tags: 
      - category
      summary: удаляет категорию
      description: Возращает удаленную категорию
      parameters:
        - name: id
          in: query
          description: id категории
          required: true
          schema:
            type: string
            default: 
      responses:
        '200':
          description: Successful operation
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Category'
        '400':
          description: Пустой параметр id
        '401':
          description: ""
        '500':
          description: Выводит ошибку
      security:
        - JwtCookieAuth: []
```
### /category/update
```
  /category/update:
    post:
      tags: 
      - category
      summary: обновляет категорию
      description: Возращает обновленную категорию
      parameters:
        - name: name
          in: query
          description: новое название категории
          required: true
          schema:
            type: string
            default:
        - name: id
          in: query
          description: id категории
          required: true
          schema:
            type: string
            default:
      responses:
        '200':
          description: Successful operation
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Category'
        '400':
          description:  Пустой параметр id
        '401':
          description: ""
        '500':
          description: Выводит ошибку
      security:
        - JwtCookieAuth: []
```

### /goods/create
```
  /goods/create:
    post:
      tags: 
      - goods
      summary: Создает категорию
      description: Возращает созданую категорию
      parameters:
        - name: name
          in: query
          description: название товара
          required: true
          schema:
            type: string
            default:
        - name: category
          in: query
          description: название категории
          required: true
          schema:
            type: string
            default:
      responses:
        '200':
          description: Successful operation
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Product'
        '400':
          description: Пустой запрос
        '401':
          description: ""
        '500':
          description: Выводит ошибку
      security:
        - JwtCookieAuth: []
```
### /goods/delete
```
  /goods/delete:
    post:
      tags: 
      - goods
      summary: удаляет товар
      description: Возращает удаленный товар
      parameters:
        - name: id
          in: query
          description: id товара
          required: true
          schema:
            type: string
            default: 
      responses:
        '200':
          description: Successful operation
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Goods'
        '400':
          description: Пустой параметр id
        '401':
          description: ""
        '500':
          description: Выводит ошибку
      security:
        - JwtCookieAuth: []
```
### /goods/update
```
 /goods/update:
    post:
      tags: 
      - goods
      summary: обновляет товар
      description: Возращает обновленный товар
      parameters:
        - name: name
          in: query
          description: новое название товара
          required: true
          schema:
            type: string
            default:
        - name: id
          in: query
          description: id товара
          required: true
          schema:
            type: string
            default:
      responses:
        '200':
          description: Successful operation
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Goods'
        '400':
          description:  Пустой параметр id
        '401':
          description: ""
        '500':
          description: Выводит ошибку
      security:
        - JwtCookieAuth: []
```

### /signin
```
 /signin:
    post:
      tags: 
      - auth
      summary: аутификация
      description: аутифицирует пользователя и устанавливает jwt и refresh токен в cookie
      parameters:
        - name: guid
          in: query
          description: индификатор пользователя
          required: true
          schema:
            type: string
            default:
      responses:
        '200':
          description: OK
          headers:
            Set-Cookie:
              description: >
                  Содержит сессионный файл cookie с именем token и refresh. Передавайте этот файл cookie обратно в последующих запросах.
              schema: 
                type: string
```

### /refresh
```
  /refresh:
    post:
      tags: 
      - auth
      summary: аутификация
      description: обновляет jwt и refresh токен в cookie удаляя старую сессию необходимо иметь refresh токен для роаботы ручки
      responses:
        '200':
          description: OK
          headers:
            Set-Cookie:
              description: >
                  Содержит сессионный файл cookie с именем token и refresh. Передавайте этот файл cookie обратно в последующих запросах.
              schema: 
                type: string
```
