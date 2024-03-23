package db

import (
	"testinhousead/internal/model"
)

func (db *sqlPostgres) AllCategories(reqID string) ([]model.Category, error) {
	var categories []model.Category

	str := ` SELECT * FROM categories; `

	row, err := db.dB.Query(str)

	if err != nil {
		db.logger.L.WithField("psql.UpdateCategories", reqID).Error("", err)
		return nil, err
	}

	var category model.Category

	for row.Next() {

		err := row.Scan(&category)
		if err != nil {
			db.logger.L.WithField("psql.Meta", reqID).Error(err)
			continue
		}
		categories = append(categories, category)
	}

	return categories, nil
}
