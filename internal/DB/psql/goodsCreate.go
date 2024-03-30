package psql

import (
	"fmt"
	"testinhousead/internal/model"
)

func (db *psql) CreateGoods(reqId string, name string) (*model.Goods, error) {
	var goods model.Goods
	var ok bool

	db.logger.L.WithField("psql.CreateGoods", reqId).Debug("полученные данные --- ", name)

	err := db.dB.QueryRow("SELECT EXISTS(SELECT * FROM goods WHERE name=$1)", name).Scan(&ok)
	if err != nil {
		return nil, err
	}

	if !ok {

		err = db.dB.QueryRow(
			`INSERT INTO goods(name) 
		 VALUES ($1) returning *`,
			name,
		).Scan(&goods.ID, &goods.Name, &goods.Removed, &goods.Updated_at, &goods.Created_at)

		db.logger.L.WithField("psql.CreateGoods", reqId).Debug("запрос --- ", goods)
		if err != nil {
			return nil, fmt.Errorf("ошибка при создании товара %s", err)
		}

		db.logger.L.WithField("psql.CreateGoods", reqId).Debug("выходные данные ---- ", goods)
		return &goods, nil
	} else {

		err = db.dB.QueryRow("SELECT * FROM goods WHERE name=$1", name).Scan(&goods.ID, &goods.Name, &goods.Removed, &goods.Updated_at, &goods.Created_at)
		if err != nil {
			return nil, fmt.Errorf("ошибка при получении товара %s", err)
		}

		db.logger.L.WithField("psql.CreateGoods", reqId).Debug("выходные данные ---- ", goods)

		return &goods, nil
	}

}
