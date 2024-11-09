package handlers

import (
	"log/slog"

	docUseCases "github.com/cfif1982/eds/internal/eds/useCases/document"
	"github.com/cfif1982/eds/internal/models"
	edsv1 "github.com/cfif1982/eds/protos/gen"
)

// Q: приходится объявлять этот интерфейс в двух местах: handlers и useCases
// Это правильно?
type DocumentRepo interface {
	Add(doc *models.Document) error
}

type Handlers struct {
	edsv1.UnimplementedEDSServer
	log         *slog.Logger
	docUseCases *docUseCases.UseCases
}

func NewHandlers(log *slog.Logger, docRepo DocumentRepo) *Handlers {
	return &Handlers{
		log:         log,
		docUseCases: docUseCases.NewUseCases(log, docRepo),
	}
}
