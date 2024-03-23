package main

import (
	"fmt"
	db "testinhousead/internal/DB"
	"testinhousead/internal/logger"

	"github.com/joho/godotenv"
)

func main() {

	log := logger.New()
	err := godotenv.Load()
	if err != nil {
		log.L.Info("Не загружается .env файл")
	}
	db, err := db.InitDb(log)
	if err != nil {
		fmt.Println(err)
	}

	sham, err := db.CreateGoods("", "Shamil")
	if err != nil {
		fmt.Println(err)
	}
	s, err := db.DeleteGoods("", 2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(sham, s)
}
