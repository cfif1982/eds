package handlers

import (
	"github.com/cfif1982/eds/pkg/logger"
)

// Интерфейс репозитория.
type DocumentRepositoryInterface interface {
}

// структура хэндлера.
type Handler struct {
	documentRepo DocumentRepositoryInterface // репозиторий для документа
	logger       *logger.Logger              // логгер
}

// создаем новый хэндлер.
func NewHandler(documentRepo DocumentRepositoryInterface, logger *logger.Logger) *Handler {
	return &Handler{
		documentRepo: documentRepo,
		logger:       logger,
	}
}
