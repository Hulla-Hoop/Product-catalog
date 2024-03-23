package db

import (
	"fmt"
	"testinhousead/internal/model"
	"time"
)

func (db *sqlPostgres) DeleteCategory(reqId string, id int) (*model.Category, error) {
	var category model.Category
	now := time.Now()
	str := fmt.Sprintf(`
	BEGIN;

	SELECT * FROM categories WHERE id = %d FOR UPDATE;

	UPDATE categories
	SET removed = true,updated_at=%s
	WHERE id=%d
	returning id,name,removed,updated_at,created_at;

	COMMIT;
	
	`, id, now, id)

	db.logger.L.WithField("psql.DeleteGOODS", reqId).Debug(" выходные данные ---", str)

	err := db.dB.QueryRow(str).Scan(&category.ID, &category.Name, &category.Removed, &category.Created_at, &category.Updated_at)

	if err != nil {
		db.logger.L.WithField("psql.DeleteGOODS", reqId).Error("", err)
		return nil, err
	}

	db.logger.L.WithField("psql.DeleteGOODS", reqId).Debug(" выходные данные ----", category)
	return &category, nil

}
