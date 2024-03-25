package handlers

import (
	"testinhousead/internal/logger"
	"testinhousead/internal/service"
)

type marketHandlers struct {
	logger  *logger.Logger
	service service.MarketService
}

func New(log *logger.Logger, service service.MarketService) *marketHandlers {
	return &marketHandlers{
		logger:  log,
		service: service,
	}
}
