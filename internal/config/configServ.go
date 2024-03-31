package config

import "os"

type configServ struct {
	Host string
	Port string
}

func ServNew() *configServ {

	return &configServ{
		Host: os.Getenv("SERVER_HOST"),
		Port: os.Getenv("SERVER_PORT"),
	}
}
