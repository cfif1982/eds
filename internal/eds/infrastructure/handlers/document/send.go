package document

import (
	"context"
	"errors"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cfif1982/eds/internal/models"
	edsv1 "github.com/cfif1982/eds/protos/gen"
	"github.com/google/uuid"
)

func (h *Handlers) SendDocument(
	ctx context.Context,
	req *edsv1.SendDocumentRequest,
) (*edsv1.SendDocumentResponse, error) {
	// Проверяем входящий ID создателя документа
	if req.GetDocumentId() == "" {
		h.log.Error("validation error", slog.String("error", "empty document id"))
		return nil, status.Error(codes.InvalidArgument, "document id required")
	}

	// Проверяем входящий  массив url файлов
	if len(req.GetFilesUrl()) == 0 {
		h.log.Error("validation error", slog.String("error", "empty files"))
		return nil, status.Error(codes.InvalidArgument, "files required")
	}

	// Проверяем входящий  массив email подписантов
	if len(req.GetSignersMail()) == 0 {
		h.log.Error("validation error", slog.String("error", "empty signers"))
		return nil, status.Error(codes.InvalidArgument, "signers required")
	}

	// конвертируем строку в uuid
	documentUUID, err := uuid.Parse(req.GetDocumentId())
	// Если возникла ошибка, то возвращаем gRPC ошибку
	if err != nil {
		h.log.Error("parse error", slog.Any("wrong document id", err))
		// Если возникла ошибка, то возвращаем код - codes.Internal
		return nil, status.Error(codes.InvalidArgument, "wrong document id")
	}

	// вызываем Сервис отправки документа
	err = h.services.SendDocument(
		ctx,
		documentUUID,
		req.GetSignersMail(),
		req.GetFilesUrl(),
	)

	if err != nil {
		// логирую ошибку в хэндлере
		h.log.Error("SendDocument() handler error", slog.Any("error", err))

		switch {
		case errors.Is(err, models.ErrDocumentNotFound):
			return nil, status.Error(codes.Internal, "document not found")
		case errors.Is(err, models.ErrUserNotFound):
			return nil, status.Error(codes.Internal, "user not found")
		default:
			return nil, status.Error(codes.Internal, "internal error")
		}
	}

	// в ответ высылаю пустую структуру, т.к. никаких данных возвращать не нужно
	result := &edsv1.SendDocumentResponse{}

	return result, nil
}
