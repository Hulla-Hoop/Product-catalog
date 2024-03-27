package config

import (
	"os"
)

type configMongo struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func MongoNew() *configMongo {

	return &configMongo{
		Host:     os.Getenv("MONGO_HOST"),
		Port:     os.Getenv("MONGO_PORT"),
		User:     os.Getenv("MONGO_USER"),
		Password: os.Getenv("MONGO_PASSWORD"),
		DBName:   os.Getenv("MONGO_NAME"),
	}
}
