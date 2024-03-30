package psql

import (
	"fmt"
	"testinhousead/internal/model"
)

func (db *psql) CreateCategory(reqId string, name string) (*model.Category, error) {
	var category model.Category
	var ok bool

	db.logger.L.WithField("psql.CreateCategory", reqId).Debug("полученные данные---", name)

	// проверка наличия записи
	err := db.dB.QueryRow("SELECT EXISTS(SELECT * FROM categories WHERE name=$1)", name).Scan(&ok)
	if err != nil {
		return nil, err
	}

	if !ok {
		//создаем новую запись
		err = db.dB.QueryRow(
			`INSERT INTO categories(name) 
		 VALUES ($1) returning *`,
			name,
		).Scan(&category.ID, &category.Name, &category.Removed, &category.Updated_at, &category.Created_at)

		if err != nil {
			return nil, fmt.Errorf("ошибка при создании товара %s", err)
		}

		db.logger.L.WithField("psql.CreateCategory", reqId).Debug("db create выходные данные ---", category)

		return &category, nil

	} else {
		//получаем существующую запись
		err = db.dB.QueryRow("SELECT * FROM categories WHERE name=$1", name).Scan(&category.ID, &category.Name, &category.Removed, &category.Updated_at, &category.Created_at)
		if err != nil {
			return nil, fmt.Errorf("ошибка при создании категории %s", err)
		}

		db.logger.L.WithField("psql.CreateCategory", reqId).Debug("db create выходные данные ---", category)

		return &category, nil
	}

}
