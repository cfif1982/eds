package documents

import (
	"context"
	"log/slog"

	edsv1 "github.com/cfif1982/eds/protos/gen"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handlers) SignDocument(
	ctx context.Context,
	req *edsv1.SignDocumentRequest,
) (*edsv1.SignDocumentResponse, error) {
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

}
