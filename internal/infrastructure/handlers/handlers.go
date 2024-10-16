package handlers

import (
	"context"
	"log/slog"

	"github.com/cfif1982/eds/internal/controller"
	docHandlers "github.com/cfif1982/eds/internal/infrastructure/handlers/document"
	edsv1 "github.com/cfif1982/eds/protos/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handlers struct {
	edsv1.UnimplementedEDSServer
	log *slog.Logger
	// для регистрации хэндлеров в grpc, этот хэндлер должен реализовывать все методы
	// но я хочу разбить их на группы для удобства
	// поэтому здесь имплементируем без имени параметра, чтобы эти методы имплементирвались напрямую
	docHandlers.DocHandlers
	// TODO: add other handlers
}

func NewHandlers(log *slog.Logger, controller *controller.Controller) *Handlers {
	return &Handlers{
		log:         log,
		DocHandlers: *docHandlers.NewHandlers(log, controller),
	}
}

// Q: если переношу этот метод в файл add.go, то возникает ошибка
// не могу вызвать edsv1.RegisterEDSServer(gRPCServer, handlers) в server.go
func (h *Handlers) AddNewDocument(
	ctx context.Context,
	req *edsv1.AddNewDocumentRequest,
) (*edsv1.AddNewDocumentResponse, error) {
	return h.AddNewDocument2(ctx, req)
}

func (h *Handlers) GetDocumentByID(
	ctx context.Context,
	req *edsv1.GetDocumentByIDRequest,
) (*edsv1.GetDocumentByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
}
