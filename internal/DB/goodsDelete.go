package db

import (
	"fmt"
	"testinhousead/internal/model"
	"time"
)

func (db *sqlPostgres) DeleteGoods(reqId string, id int) (*model.Goods, error) {
	var goods model.Goods
	now := time.Now()
	str := fmt.Sprintf(`
	BEGIN;

	SELECT * FROM goods WHERE id = %d FOR UPDATE;

	UPDATE goods 
	SET removed = true,updated_at=%s
	WHERE id=%d
	returning id,name,removed,updated_at,created_at;

	COMMIT;
	
	`, id, now, id)

	db.logger.L.WithField("psql.DeleteGOODS", reqId).Debug(" выходные данные ---", str)

	err := db.dB.QueryRow(str).Scan(&goods.ID, &goods.Name, &goods.Removed, &goods.Created_at, &goods.Updated_at)

	if err != nil {
		db.logger.L.WithField("psql.DeleteGOODS", reqId).Error("", err)
		return nil, err
	}

	db.logger.L.WithField("psql.DeleteGOODS", reqId).Debug(" выходные данные ----", goods)
	return &goods, nil

}
