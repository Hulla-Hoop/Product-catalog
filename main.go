package main

import (
	"testinhousead/internal/logger"
	"testinhousead/pkg/app"

	"github.com/joho/godotenv"
)

func main() {

	log := logger.New()
	err := godotenv.Load()
	if err != nil {
		log.L.Info("Не загружается .env файл")
	}
	a, err := app.NewApp()
	if err != nil {
		panic(err)
	}

	a.Start()

	/* parser.Parse() */
}
