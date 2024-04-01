package app

import (
	"net/http"
	"testinhousead/internal/DB/mongo"
	"testinhousead/internal/DB/psql"
	"testinhousead/internal/config"
	"testinhousead/internal/handlers"
	aut "testinhousead/internal/handlers/autification"
	hand "testinhousead/internal/handlers/catalog"
	"testinhousead/internal/handlers/middlware"
	"testinhousead/internal/logger"
	"testinhousead/internal/service/autification"
	"testinhousead/internal/service/catalog"
)

type first struct {
	hand handlers.Catalog
}

type second struct {
	hand handlers.Aut
}

type app struct {
	first  *first
	second *second
	mux    *http.ServeMux
	logger *logger.Logger
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

	//обьявляем второй сервис и подключаем зависимости
	second := second{}
	mongo := mongo.New(l)
	se := autification.NewAut(l, mongo)
	second.hand = aut.NewAut(l, se)

	// создаем mux
	mux := http.NewServeMux()

	// обьявление ручек. middleware.ReqID прокидывает реквест айди для удобства просмотра логов
	// middlware.Aut проверят пользователей на наличие прав
	mux.HandleFunc("/allcategories/", middlware.ReqID(first.hand.AllCategories))
	mux.HandleFunc("/goodsoncategory", middlware.ReqID(first.hand.GoodsOnCateory))
	mux.HandleFunc("/category/create", middlware.ReqID(middlware.Aut(first.hand.CreateCategory)))
	mux.HandleFunc("/category/delete", middlware.ReqID(middlware.Aut(first.hand.DeleteCategory)))
	mux.HandleFunc("/category/update", middlware.ReqID(middlware.Aut(first.hand.UpdateCategory)))
	mux.HandleFunc("/goods/create", middlware.ReqID(middlware.Aut(first.hand.CreateGoods)))
	mux.HandleFunc("/goods/delete", middlware.ReqID(middlware.Aut(first.hand.DeleteGoods)))
	mux.HandleFunc("/goods/update", middlware.ReqID(middlware.Aut(first.hand.UpdateGoods)))
	// ручки для входа и обновления токена доступа
	mux.HandleFunc("/signin", middlware.ReqID(second.hand.SignIn))
	mux.HandleFunc("/refresh", middlware.ReqID(second.hand.Refresh))

	return &app{
		first:  &first,
		second: &second,
		mux:    mux,
		logger: l,
	}, nil
}

// функция запуска приложения
func (a *app) Start() {
	conf := config.ServNew()

	err := http.ListenAndServe(conf.Host+":"+conf.Port, a.mux)

	if err != nil {
		a.logger.L.WithField("APP.Start", err).Error()
	}

	a.logger.L.WithField("APP.Start", "").Infof("Сервер стартовал на %s:%s", conf.Host, conf.Port)
}
