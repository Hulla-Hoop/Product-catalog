package psql

import (
	"fmt"
	"testinhousead/internal/model"
	"time"
)

func (db *psql) DeleteCategory(reqId string, id int) (*model.Category, error) {
	var category model.Category
	now := time.Now().Format(time.DateTime)
	str := fmt.Sprintf(`
	BEGIN;

	UPDATE categories
	SET removed = true,updated_at='%s'
	WHERE id=%d
	returning id,name,removed,updated_at,created_at;

	SELECT * FROM categories WHERE id = %d FOR UPDATE;

	COMMIT;
	
	`, now, id, id)

	db.logger.L.WithField("psql.DeleteCategory", reqId).Debug(" Строка запроса ---", str)

	err := db.dB.QueryRow(str).Scan(&category.ID, &category.Name, &category.Removed, &category.Created_at, &category.Updated_at)

	if err != nil {
		db.logger.L.WithField("psql.DeleteCategory", reqId).Error("", err)
		return nil, err
	}

	db.logger.L.WithField("psql.DeleteCategory", reqId).Debug(" выходные данные ----", category)
	return &category, nil

}
