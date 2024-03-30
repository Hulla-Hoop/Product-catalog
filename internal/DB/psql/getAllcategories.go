package psql

import (
	"testinhousead/internal/model"
)

func (db *psql) AllCategories(reqID string) ([]model.Category, error) {
	var categories []model.Category

	str := ` SELECT * FROM categories WHERE removed=false; `

	row, err := db.dB.Query(str)

	if err != nil {
		db.logger.L.WithField("psql.AllCategories", reqID).Error("", err)
		return nil, err
	}

	var category model.Category

	for row.Next() {

		err := row.Scan(&category.ID, &category.Name, &category.Removed, &category.Updated_at, &category.Created_at)
		if err != nil {
			db.logger.L.WithField("psql.AllCategories", reqID).Error(err)
			continue
		}
		categories = append(categories, category)
	}

	db.logger.L.WithField("psql.AllCategories", reqID).Debug("выходные данные ---- ", categories)

	return categories, nil
}
