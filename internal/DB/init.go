package db

import (
	"database/sql"
	"fmt"
	"testinhousead/internal/config"
	"testinhousead/internal/logger"

	_ "github.com/lib/pq"
)

type sqlPostgres struct {
	dB     *sql.DB
	logger *logger.Logger
}

func InitDb(logger *logger.Logger) (*sqlPostgres, error) {
	config := config.DbNew()

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=%s", config.Host, config.User, config.DBName, config.Password, config.Port, config.SSLMode)
	logger.L.Info(dsn)
	dB, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	/* err = goose.Up(dB, "migrations/psql")
	if err != nil {
		return nil,
			fmt.Errorf("--- Ошибка миграции:%s", err)
	} */
	return &sqlPostgres{
		dB:     dB,
		logger: logger,
	}, nil

}

func (s *sqlPostgres) Close() error {
	return s.dB.Close()
}
