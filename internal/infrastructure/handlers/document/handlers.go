package document

import (
	"log/slog"

	"github.com/cfif1982/eds/internal/controller"
)

type DocHandlers struct {
	log        *slog.Logger
	controller *controller.Controller
}

func NewHandlers(log *slog.Logger, controller *controller.Controller) *DocHandlers {
	return &DocHandlers{
		log:        log,
		controller: controller,
	}
}
