package handlers

import "net/http"

type Aut interface {
	SignIn(w http.ResponseWriter, r *http.Request)
	Refresh(w http.ResponseWriter, r *http.Request)
}

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
