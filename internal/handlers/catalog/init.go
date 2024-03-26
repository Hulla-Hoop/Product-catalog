package handlers

import (
	"testinhousead/internal/logger"
	"testinhousead/internal/service"
)

type marketHandlers struct {
	logger  *logger.Logger
	service service.Cataloger
}

func NewCatalog(log *logger.Logger, service service.Cataloger) *marketHandlers {
	return &marketHandlers{
		logger:  log,
		service: service,
	}
}
