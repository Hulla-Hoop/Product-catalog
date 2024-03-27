package autification

import (
	"testinhousead/internal/DB/mongo"
	"testinhousead/internal/logger"
)

type jWT struct {
	logger *logger.Logger
	db     *mongo.Mongo
}

func NewAut(log *logger.Logger, db *mongo.Mongo) *jWT {
	return &jWT{
		logger: log,
		db:     db,
	}
}
