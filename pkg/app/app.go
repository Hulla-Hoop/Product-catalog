package app

import (
	"fmt"
	"net/http"
	"testinhousead/internal/DB/mongo"
	"testinhousead/internal/DB/psql"
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
}

func NewApp() (*app, error) {
	first := first{}
	second := second{}

	l := logger.New()
	db, err := psql.InitDb(l)
	if err != nil {
		return nil, err
	}
	s := catalog.NewCatalog(l, db)
	first.hand = hand.NewCatalog(l, s)

	mongo := mongo.New(l)

	se := autification.NewAut(l, mongo)

	second.hand = aut.NewAut(l, se)

	mux := http.NewServeMux()

	mux.HandleFunc("/allcategories/", first.hand.AllCategories)
	mux.HandleFunc("/goodsoncategory", first.hand.GoodsOnCateory)
	mux.HandleFunc("/category/create", middlware.ReqID(middlware.Aut(first.hand.CreateCategory)))
	mux.HandleFunc("/category/delete", middlware.ReqID(middlware.Aut(first.hand.DeleteCategory)))
	mux.HandleFunc("/category/update", middlware.ReqID(middlware.Aut(first.hand.UpdateCategory)))
	mux.HandleFunc("/goods/create", middlware.ReqID(middlware.Aut(first.hand.CreateGoods)))
	mux.HandleFunc("/goods/delete", middlware.ReqID(middlware.Aut(first.hand.DeleteGoods)))
	mux.HandleFunc("/goods/update", middlware.ReqID(middlware.Aut(first.hand.UpdateGoods)))

	mux.HandleFunc("/signin", middlware.ReqID(second.hand.SignIn))
	mux.HandleFunc("/refresh", middlware.ReqID(second.hand.Refresh))

	return &app{
		first:  &first,
		second: &second,
		mux:    mux,
	}, nil
}

func (a *app) Start() {
	err := http.ListenAndServe(":8080", a.mux)
	if err != nil {
		fmt.Println(err)
	}
}
