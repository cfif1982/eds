package controller

import (
	"log/slog"

	docUseCases "github.com/cfif1982/eds/internal/eds/useCases/document"
	"github.com/cfif1982/eds/internal/models"
)

// Q: приходится объявлять этот интерфейс в двух местах: controller и useCases
// Это правильно?
type DocumentRepo interface {
	Add(doc *models.Document) error
}

type Controller struct {
	log         *slog.Logger
	docUseCases *docUseCases.UseCases
}

func NewController(log *slog.Logger, docRepo DocumentRepo) *Controller {
	return &Controller{
		log:         log,
		docUseCases: docUseCases.NewUseCases(log, docRepo),
	}
}
