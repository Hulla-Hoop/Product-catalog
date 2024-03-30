package psql

import (
	"fmt"
	"testinhousead/internal/model"
	"time"
)

func (db *psql) UpdateGoods(reqId string, id int, name string) (*model.Goods, error) {

	db.logger.L.WithField("psql.UpdateGOODS", reqId).Debug(" выходные данные ----", id, name)

	var goods model.Goods

	now := time.Now().Format(time.DateTime)

	str := fmt.Sprintf(`
	BEGIN;


	UPDATE goods 
	SET updated_at='%s',name='%s'
	WHERE id=%d
	returning id,name,removed,updated_at,created_at;

	SELECT * FROM goods WHERE id = %d ;

	COMMIT;
	
	`, now, name, id, id)

	db.logger.L.WithField("psql.UpdateGOODS", reqId).Debug(" запрос ---", str)

	err := db.dB.QueryRow(str).Scan(&goods.ID, &goods.Name, &goods.Removed, &goods.Created_at, &goods.Updated_at)

	if err != nil {
		db.logger.L.WithField("psql.UpdateGOODS", reqId).Error("", err)
		return nil, err
	}

	db.logger.L.WithField("psql.UpdateGOODS", reqId).Debug(" выходные данные ----", goods)
	return &goods, nil

}
