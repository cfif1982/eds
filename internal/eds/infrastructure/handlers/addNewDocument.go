package handlers

import (
	"context"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	edsv1 "github.com/cfif1982/eds/protos/gen"
	"github.com/google/uuid"
)

// Q: так норм делать?
const (
	emptyValueStr = "" // констанат для пустого значения у string. Для удобства чтения кода
)

func (h *Handlers) AddNewDocument(
	ctx context.Context,
	req *edsv1.AddNewDocumentRequest,
) (*edsv1.AddNewDocumentResponse, error) {
	// Q: работа с ошибками.
	// Проверяем входящий ID создателя документа
	if req.GetCreatorId() == emptyValueStr {
		h.log.Error("validation error", slog.String("error", "empty creator id"))
		return nil, status.Error(codes.InvalidArgument, "creator id required")
	}

	// конвертируем строку в uuid
	creatorUUID, err := uuid.Parse(req.GetCreatorId())
	// Q: работа с ошибками.
	// Если возникла ошибка, то возвращаем gRPC ошибку
	if err != nil {
		h.log.Error("parse error", slog.Any("wrong creator id", err))
		// Если возникла ошибка, то возвращаем код - codes.Internal
		return nil, status.Error(codes.InvalidArgument, "wrong creator id")
	}

	// вызываем UseCase добавления документа
	qrCode, err := h.docUseCases.Add(ctx, creatorUUID)

	// Q: работа с ошибками
	if err != nil {
		// логирую ошибку в хэндлере
		h.log.Error("add document useCase error", slog.Any("error", err))
		// здесь проверяем ошибку. Если это одна из моих ошибок из models,
		// то выводить нужные коды ошибок
		// if errors.Is(err, models.ErrInvalidUser) || errors.Is(err, models.ErrUserValidation) else
		// Если возникла ошибка, то возвращаем код - codes.Internal
		return nil, status.Error(codes.Internal, "internal error")
	}

	// конвертируем данные из domain в grpc для ответа.
	result := &edsv1.AddNewDocumentResponse{
		QrCode: qrCode,
	}

	return result, nil
}
