package documents

import (
	"context"
	"errors"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cfif1982/eds/internal/models"
	edsv1 "github.com/cfif1982/eds/protos/gen"
)

func (h *Handlers) AddNewDocument(
	ctx context.Context,
	req *edsv1.AddNewDocumentRequest,
) (*edsv1.AddNewDocumentResponse, error) {
	// Проверяем входящий ID создателя документа
	if req.GetCreatorEmail() == "" {
		h.log.Error("validation error", slog.String("error", "empty creator email"))
		return nil, status.Error(codes.InvalidArgument, "creator email required")
	}

	// TODO сделать валидацию email
	/*
		err := validateEmail(req.GetCreatorEmail())
		// Если возникла ошибка, то возвращаем gRPC ошибку
		if err != nil {
			h.log.Error("validation error", slog.Any("wrong creator email", err))
			// Если возникла ошибка, то возвращаем код - codes.Internal
			return nil, status.Error(codes.InvalidArgument, "wrong creator email")
		}
	*/

	// вызываем сервис добавления документа
	qrCode, err := h.services.AddDocument(ctx, req.GetCreatorEmail())

	if err != nil {
		// логирую ошибку в хэндлере
		h.log.Error("AddNewDocument() handler error", slog.Any("error", err))

		// Q: так нужно возвращать ошибки юзеру?
		switch {
		case errors.Is(err, models.ErrDocumentAlreadyExists):
			return nil, status.Error(codes.Internal, "document already exist")
		case errors.Is(err, models.ErrUserNotFound):
			return nil, status.Error(codes.Internal, "user not found")
		default:
			return nil, status.Error(codes.Internal, "internal error")
		}
	}

	// конвертируем данные из domain в grpc для ответа.
	result := &edsv1.AddNewDocumentResponse{
		QrCode: qrCode,
	}

	return result, nil
}
