package main

import (
	"fmt"
	"net/http"
	"testinhousead/internal/DB/psql"
	handlers "testinhousead/internal/handlers/goods"
	"testinhousead/internal/logger"
	"testinhousead/internal/service/goods"

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
	s := goods.New(log, db)
	h := handlers.New(log, s)

	http.HandleFunc("/", h.CreateCategory)

	http.ListenAndServe(":8080", nil)

	cat, err := db.GoodsOnCateory("", "Flour")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cat)

}
