package main

import (
	"fmt"
	"net/http"
	"testinhousead/internal/DB/psql"
	handlers "testinhousead/internal/handlers/goods"
	"testinhousead/internal/logger"
	service "testinhousead/internal/service/goods"

	"github.com/joho/godotenv"
)

func main() {

	log := logger.New()
	err := godotenv.Load()
	if err != nil {
		log.L.Info("Не загружается .env файл")
	}
	db, err := psql.InitDb(log)
	if err != nil {
		fmt.Println(err)
	}
	s := service.New(log, db)
	h := handlers.New(log, s)

	http.HandleFunc("/category/create", h.CreateCategory)
	http.HandleFunc("/category/delete", h.DeleteCategory)
	http.HandleFunc("/category/update", h.UpdateCategory)
	http.HandleFunc("/goods/create", h.CreateGoods)
	http.HandleFunc("/goods/delete", h.DeleteGoods)
	http.HandleFunc("/goods/update", h.UpdateGoods)
	http.HandleFunc("/allcategories", h.AllCategories)
	http.HandleFunc("/goodsoncategory", h.GoodsOnCateory)

	http.ListenAndServe(":8080", nil)

	cat, err := db.GoodsOnCateory("", "Flour")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cat)

}
