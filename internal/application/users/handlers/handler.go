package handlers

import (
	"github.com/cfif1982/eds/pkg/logger"
)

// Интерфейс репозитория.
type UserRepositoryInterface interface {
}

// структура хэндлера.
type Handler struct {
	userRepo UserRepositoryInterface // репозиторий для документа
	logger   *logger.Logger          // логгер
}

// создаем новый хэндлер.
func NewHandler(userRepo UserRepositoryInterface, logger *logger.Logger) *Handler {
	return &Handler{
		userRepo: userRepo,
		logger:   logger,
	}
}
