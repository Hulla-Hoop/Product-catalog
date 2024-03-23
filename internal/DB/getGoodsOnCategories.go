package db

import (
	"fmt"
	"testinhousead/internal/model"
)

func (db *sqlPostgres) GoodsOnCateory(reqID string, category string) ([]model.Product, error) {
	var goods []model.Product

	str := fmt.Sprintf(`
	SELECT goods.name,categories.name FROM goods
	JOIN relation
	ON goods.id=relation.goods_id
	JOIN categories
	ON categories.id = relation.category_id
	WHERE categories.name=%s;
	`, category)

	row, err := db.dB.Query(str)

	if err != nil {
		db.logger.L.WithField("psql.UpdateCategories", reqID).Error("", err)
		return nil, err
	}

	var product model.Product

	for row.Next() {

		err := row.Scan(&product.Name, &product.Category)
		if err != nil {
			db.logger.L.WithField("psql.Meta", reqID).Error(err)
			continue
		}
		goods = append(goods, product)
	}

	return goods, nil

}
