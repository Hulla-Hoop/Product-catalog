package db

import (
	"database/sql"
	"fmt"
	"testinhousead/internal/model"
)

func (db *sqlPostgres) CreateGoods(reqId string, name string, project_id int) (*model.Goods, error) {
	var goods model.Goods

	db.logger.L.WithField("psql.Create", reqId).Debug("db create полученные данные---", name, "----", project_id)
	err := db.dB.QueryRow(
		`INSERT INTO goods(name,categories_id) 
		 VALUES ($1,$2) returning id,name,categories_id`,
		name,
		project_id,
	).Scan(&goods.ID, &goods.Name, &goods.CategryID)

	db.logger.L.WithField("psql.Create", reqId).Debug("db create выходные данные ---", goods)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, fmt.Errorf("товар добавлен но не удалось записать ID %s", err)
		default:
			return nil, fmt.Errorf("ошибка при создании товара %s", err)
		}
	}

	db.logger.L.WithField("psql.Create", reqId).Debug("db create выходные данные ----", goods)
	return &goods, nil

}
