package document

import (
	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cfif1982/eds/internal/models"
	edsv1 "github.com/cfif1982/eds/protos/gen"
)

func (h *Handlers) GetDocumentByID(
	ctx context.Context,
	req *edsv1.GetDocumentByIDRequest,
) (*edsv1.GetDocumentByIDResponse, error) {
	// Проверяем входящий ID  документа
	if req.GetDocumentId() == "" {
		h.log.Error("validation error", slog.String("error", "empty document id"))
		return nil, status.Error(codes.InvalidArgument, "document id required")
	}

	// конвертируем строку в uuid
	documentUUID, err := uuid.Parse(req.GetDocumentId())
	// Если возникла ошибка, то возвращаем gRPC ошибку
	if err != nil {
		h.log.Error("parse error", slog.Any("wrong document id", err))
		// Если возникла ошибка, то возвращаем код - codes.Internal
		return nil, status.Error(codes.InvalidArgument, "wrong document id")
	}

	// вызываем Сервис получения документа
	doc, err := h.services.GetById(ctx, documentUUID)

	if err != nil {
		// логирую ошибку в хэндлере
		h.log.Error("GetDocumentByID() error", slog.Any("error", err))

		switch {
		case errors.Is(err, models.ErrDocumentNotFound):
			return nil, status.Error(codes.Internal, "document not found")
		default:
			return nil, status.Error(codes.Internal, "internal error")
		}
	}

	// конвертируем данные из сервиса в grpc для ответа.
	result := h.documentToGetDocumentByIDResponse(doc)

	return result, nil
}
