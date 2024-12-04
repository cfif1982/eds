package documents

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

		// Q: так нужно возвращать ошибки юзеру?
		switch {
		case errors.Is(err, models.ErrDocumentNotFound):
			return nil, status.Error(codes.Internal, "document not found")
		default:
			return nil, status.Error(codes.Internal, "internal error")
		}
	}

	// Q: как тут конвертировать? откуда брать инфу по юзерам?
	// конвертируем полученную модель документа в модель grpc
	// получаем модель user для создалтеля документа
	// creator := h.services.

	// получаем слайс моделей user для подписантов документа
	signers := make([]*models.User, 0, len(doc.Signers))
	for _, id := range doc.Signers {
		// signer, err := h.services.
		signers = append(signers, signer)
	}

	// получаем слайс моделей file для файлов документа
	files := make([]*models.File, 0, len(doc.Files))
	for _, id := range doc.Files {
		// file, err := h.services.
		files = append(files, file)
	}

	// конвертируем данные из сервиса в grpc для ответа.
	result := &edsv1.GetDocumentByIDResponse{
		DocumentId: doc.ID.String(),
		Creator:    creator,
		Signers:    signers,
		FilesUrl:   files,
	}

	return result, nil
}
