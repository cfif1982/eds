package handlers

import (
	"log/slog"

	docUseCases "github.com/cfif1982/eds/internal/eds/useCases/document"
	edsv1 "github.com/cfif1982/eds/protos/gen"
)

type Handlers struct {
	edsv1.UnimplementedEDSServer
	log         *slog.Logger
	docUseCases *docUseCases.UseCases
}

func NewHandlers(log *slog.Logger, docUseCases *docUseCases.UseCases) *Handlers {
	return &Handlers{
		log:         log,
		docUseCases: docUseCases,
	}
}
