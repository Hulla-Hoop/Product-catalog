package psql

import (
	"fmt"
	"testinhousead/internal/model"
)

func (db *psql) CreateCategory(reqId string, name string) (*model.Category, error) {
	var category model.Category

	db.logger.L.WithField("psql.CreateCategory", reqId).Debug("полученные данные---", name)
	err := db.dB.QueryRow(
		`INSERT INTO categories(name) 
		 VALUES ($1) returning *`,
		name,
	).Scan(&category.ID, &category.Name, &category.Removed, &category.Updated_at, &category.Created_at)

	db.logger.L.WithField("psql.CreateCategory", reqId).Debug("db create выходные данные ---", category)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании товара %s", err)
	}

	return &category, nil

}
