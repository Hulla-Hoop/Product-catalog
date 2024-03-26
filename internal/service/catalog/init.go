package catalog

import (
	db "testinhousead/internal/DB"
	"testinhousead/internal/logger"
)

type catalog struct {
	logger *logger.Logger
	db     db.DB
}

func NewCatalog(log *logger.Logger, db db.DB) *catalog {
	return &catalog{
		logger: log,
		db:     db,
	}
}

func (s *catalog) Close() error {
	return s.db.Close()
}
