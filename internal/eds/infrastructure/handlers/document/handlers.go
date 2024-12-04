package documents

import (
	"log/slog"

	docServices "github.com/cfif1982/eds/internal/eds/services/document"
	edsv1 "github.com/cfif1982/eds/protos/gen"
)

type Handlers struct {
	edsv1.UnimplementedEDSServer
	log      *slog.Logger
	services *docServices.Services
}

func NewHandlers(log *slog.Logger, services *docServices.Services) *Handlers {
	return &Handlers{
		log:      log,
		services: services,
	}
}
