package main

import (
	"fmt"
	"testinhousead/internal/DB/psql"
	"testinhousead/internal/logger"

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

	cat, err := db.GoodsOnCateory("", "Flour")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cat)

}
