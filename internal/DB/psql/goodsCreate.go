package psql

import (
	"fmt"
	"testinhousead/internal/model"
)

func (db *psql) CreateGoods(reqId string, name string) (*model.Goods, error) {
	var goods model.Goods

	db.logger.L.WithField("psql.CreateGoods", reqId).Debug("полученные данные --- ", name)
	err := db.dB.QueryRow(
		`INSERT INTO goods(name) 
		 VALUES ($1) returning id,name`,
		name,
	).Scan(&goods.ID, &goods.Name)

	db.logger.L.WithField("psql.CreateGoods", reqId).Debug("запрос --- ", goods)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании товара %s", err)
	}

	db.logger.L.WithField("psql.CreateGoods", reqId).Debug("выходные данные ---- ", goods)
	return &goods, nil

}
