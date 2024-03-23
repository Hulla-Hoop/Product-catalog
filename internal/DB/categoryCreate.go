package db

import (
	"fmt"
	"testinhousead/internal/model"
)

func (db *sqlPostgres) CreateCategory(reqId string, name string) (*model.Category, error) {
	var category model.Category

	db.logger.L.WithField("psql.Create", reqId).Debug("db create полученные данные---", name)
	err := db.dB.QueryRow(
		`INSERT INTO category(name) 
		 VALUES ($1) returning *`,
		name,
	).Scan(&category.ID, &category.Name, &category.Updated_at, &category.Created_at)

	db.logger.L.WithField("psql.Create", reqId).Debug("db create выходные данные ---", category)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании товара %s", err)
	}

	db.logger.L.WithField("psql.Create", reqId).Debug("db create выходные данные ----", category)
	return &category, nil

}
