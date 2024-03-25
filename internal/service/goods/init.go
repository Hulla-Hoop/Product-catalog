package goods

import (
	db "testinhousead/internal/DB"
	"testinhousead/internal/logger"
)

type service struct {
	logger *logger.Logger
	db     db.DB
}

func New(log *logger.Logger, db db.DB) *service {
	return &service{
		logger: log,
		db:     db,
	}
}

func (s *service) Close() error {
	return s.db.Close()
}
