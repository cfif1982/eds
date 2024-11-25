package handlers

import (
	"context"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	edsv1 "github.com/cfif1982/eds/protos/gen"
	"github.com/google/uuid"
)

func (h *Handlers) SendDocument(
	ctx context.Context,
	req *edsv1.SendDocumentRequest,
) (*edsv1.SendDocumentResponse, error) {
	// Q: работа с ошибками.
	// Проверяем входящий ID создателя документа
	if req.GetDocumentId() == emptyValueStr {
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
	// Q: работа с ошибками.
	// Если возникла ошибка, то возвращаем gRPC ошибку
	if err != nil {
		h.log.Error("parse error", slog.Any("wrong document id", err))
		// Если возникла ошибка, то возвращаем код - codes.Internal
		return nil, status.Error(codes.InvalidArgument, "wrong document id")
	}

	// вызываем UseCase отправки документа
	err = h.docUseCases.Send(
		ctx,
		documentUUID,
		req.GetSignersMail(),
		req.GetFilesUrl(),
	)

	// Q: работа с ошибками
	if err != nil {
		// логирую ошибку в хэндлере
		h.log.Error("send document useCase error", slog.Any("error", err))
		// здесь проверяем ошибку. Если это одна из моих ошибок из models,
		// то выводить нужные коды ошибок
		// if errors.Is(err, models.ErrInvalidUser) || errors.Is(err, models.ErrUserValidation) else
		// Если возникла ошибка, то возвращаем код - codes.Internal
		return nil, status.Error(codes.Internal, "internal error")
	}

	// конвертируем данные из domain в grpc для ответа.
	// Q: правильный ответ даю? или достаточно было просто вернуть ошибку как nil?
	result := &edsv1.SendDocumentResponse{
		Success: true,
	}

	return result, nil
}
