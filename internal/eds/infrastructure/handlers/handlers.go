package handlers

import (
	"log/slog"

	"github.com/cfif1982/eds/internal/eds/controller"
	edsv1 "github.com/cfif1982/eds/protos/gen"
)

type Handlers struct {
	edsv1.UnimplementedEDSServer
	log        *slog.Logger
	controller *controller.Controller
}

func NewHandlers(log *slog.Logger, controller *controller.Controller) *Handlers {
	return &Handlers{
		log:        log,
		controller: controller,
	}
}
