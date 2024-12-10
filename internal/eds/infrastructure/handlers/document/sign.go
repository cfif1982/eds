package document

import (
	"context"
	"errors"
	"log/slog"

	"github.com/cfif1982/eds/internal/models"
	edsv1 "github.com/cfif1982/eds/protos/gen"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handlers) SignDocument(
	ctx context.Context,
	req *edsv1.SignDocumentRequest,
) (*edsv1.SignDocumentResponse, error) {
	// Проверяем входящий  url файла
	if req.GetSignatureFileUrl() == "" {
		h.log.Error("validation error", slog.String("error", "empty signature file"))
		return nil, status.Error(codes.InvalidArgument, "signature file required")
	}

	// Проверяем входящий  email подписанта
	if req.GetSignerMail() == "" {
		h.log.Error("validation error", slog.String("error", "empty signer email"))
		return nil, status.Error(codes.InvalidArgument, "signer email required")
	}

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

	// вызываем Сервис подписания документа
	err = h.services.SignDocument(ctx, documentUUID, req.GetSignerMail(), req.GetSignatureFileUrl())

	if err != nil {
		// логирую ошибку в хэндлере
		h.log.Error("SignDocument() error", slog.Any("error", err))

		switch {
		case errors.Is(err, models.ErrDocumentNotFound):
			return nil, status.Error(codes.Internal, "document not found")
		default:
			return nil, status.Error(codes.Internal, "internal error")
		}
	}

	// конвертируем данные из сервиса в grpc для ответа.
	result := &edsv1.SignDocumentResponse{}

	return result, nil
}
