package psql

import (
	"fmt"
	"testinhousead/internal/model"
)

func (db *psql) CreateRelation(reqId string, goods_id int, category_id int) (*model.Product, error) {
	var product model.Product

	db.logger.L.WithField("psql.CreateRelation", reqId).Debug("полученные данные---", "goods_id=", goods_id, "category_id=", category_id)

	str := fmt.Sprintf(`
	BEGIN;

	INSERT INTO relation(goods_id,category_id) 
	VALUES (%d,%d);

	SELECT goods.name,categories.name FROM goods
	JOIN relation
	ON goods.id=relation.goods_id
	JOIN categories
	ON categories.id = relation.category_id
	WHERE categories.id=%d AND goods.id=%d AND categories.removed=false AND goods.removed=false;

	COMMIT;`, goods_id, category_id, category_id, goods_id)

	err := db.dB.QueryRow(str).Scan(&product.Name, &product.Category)

	db.logger.L.WithField("psql.CreateRelation", reqId).Debug("выходные данные --- ", product)

	if err != nil {
		return nil, fmt.Errorf("ошибка при создании товара %s", err)
	}

	return &product, nil

}
