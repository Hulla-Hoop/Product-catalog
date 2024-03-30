package psql

import (
	"fmt"
	"testinhousead/internal/model"
)

func (db *psql) GoodsOnCateory(reqID string, category string) ([]model.Product, error) {
	var goods []model.Product

	str := fmt.Sprintf(`
	SELECT goods.name,categories.name FROM goods
	JOIN relation
	ON goods.id=relation.goods_id
	JOIN categories
	ON categories.id = relation.category_id
	WHERE categories.name='%s' AND categories.removed=false AND goods.removed=false;
	`, category)

	db.logger.L.WithField("psql.GoodsOnCateory", reqID).Debug("Query ---- ", str)

	row, err := db.dB.Query(str)

	if err != nil {
		db.logger.L.WithField("psql.GoodsOnCateory", reqID).Error("", err)
		return nil, err
	}

	var product model.Product

	for row.Next() {

		err := row.Scan(&product.Name, &product.Category)
		if err != nil {
			db.logger.L.WithField("psql.GoodsOnCateory", reqID).Error(err)
			continue
		}
		goods = append(goods, product)
	}

	db.logger.L.WithField("psql.GoodsOnCateory", reqID).Debug("Выходные данные --- ", goods)

	return goods, nil

}
