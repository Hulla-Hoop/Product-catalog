package psql

import (
	"fmt"
	"testinhousead/internal/model"
	"time"
)

func (db *psql) UpdateCategory(reqId string, id int, name string) (*model.Category, error) {
	var category model.Category
	now := time.Now().Format(time.DateTime)
	str := fmt.Sprintf(`
	BEGIN;


	UPDATE categories 
	SET updated_at='%s',name='%s'
	WHERE id=%d
	returning id,name,removed,updated_at,created_at;

	SELECT * FROM categories WHERE id = %d;

	COMMIT;
	
	`, now, name, id, id)

	db.logger.L.WithField("psql.UpdateCategories", reqId).Debug(" запрос ---", str)

	err := db.dB.QueryRow(str).Scan(&category.ID, &category.Name, &category.Removed, &category.Created_at, &category.Updated_at)

	if err != nil {
		db.logger.L.WithField("psql.UpdateCategories", reqId).Error("", err)
		return nil, err
	}

	db.logger.L.WithField("psql.UpdateCategories", reqId).Debug(" выходные данные ----", category)
	return &category, nil

}
