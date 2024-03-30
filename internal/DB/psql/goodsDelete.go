package psql

import (
	"fmt"
	"testinhousead/internal/model"
	"time"
)

func (db *psql) DeleteGoods(reqId string, id int) (*model.Goods, error) {
	var goods model.Goods
	now := time.Now().Format(time.DateTime)

	str := fmt.Sprintf(`
	BEGIN;


	UPDATE goods 
	SET updated_at='%s',removed = true
	WHERE id=%d
	returning id,name,removed,updated_at,created_at;
	SELECT * FROM goods WHERE id = %d ;
	COMMIT;
	 
	
	`, now, id, id)

	db.logger.L.WithField("psql.DeleteGOODS", reqId).Debug(" запрос --- ", str)

	err := db.dB.QueryRow(str).Scan(&goods.ID, &goods.Name, &goods.Removed, &goods.Updated_at, &goods.Created_at)

	if err != nil {
		db.logger.L.WithField("psql.DeleteGOODS", reqId).Error("", err)
		return nil, err
	}

	db.logger.L.WithField("psql.DeleteGOODS", reqId).Debug(" выходные данные ---- ", goods)
	return &goods, nil

}
