package psql

import (
	"database/sql"
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

	db.logger.L.WithField("psql.UpdateCategories", reqID).Debug("Query ---- ", str)

	row, err := db.dB.Query(str)

	if err != nil {
		db.logger.L.WithField("psql.UpdateCategories", reqID).Error("", err)
		return nil, err
	}

	var product model.Product

	for row.Next() {

		err := row.Scan(&product.Name, &product.Category)
		if err != nil {
			if err == sql.ErrNoRows {
				db.logger.L.WithField("psql.Meta", reqID).Error(err)
				continue
			} else {
				db.logger.L.WithField("psql.Meta", reqID).Error(err)
				return nil, err
			}
		}
		goods = append(goods, product)
	}

	return goods, nil

}
