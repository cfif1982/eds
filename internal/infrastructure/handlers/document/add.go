package document

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	edsv1 "github.com/cfif1982/eds/protos/gen"
	"github.com/google/uuid"
)

const (
	emptyValueInt = "" // констанат для пустого значения у int. Для удобства чтения кода
)

func (h *DocHandlers) AddNewDocument2(
	ctx context.Context,
	req *edsv1.AddNewDocumentRequest,
) (*edsv1.AddNewDocumentResponse, error) {
	// проверяем входящие данные
	err := validateAddNewDocument(req)
	if err != nil {
		return nil, err
	}

	// вызываем UseCase добавления документа
	result, err := h.controller.AddNewDocument(req.GetCreatorId())
	if err != nil {
		// TODO: тут нужно доабваить более широкую обработку ошибки
		// Например, такой пользователь уже существует и т.д.
		// здесь конкретно это не нужно, но на будущее себе пометил

		// Если возникла ошибка, то возвращаем код - codes.Internal
		return nil, status.Error(codes.Internal, "internal error")
	}

	return result, nil
}

// проверяем входящие данные.
func validateAddNewDocument(req *edsv1.AddNewDocumentRequest) error {
	// Проверяем входящий ID создателя документа
	// проверку делаем через геттер - Get...
	if req.GetCreatorId() == emptyValueInt {
		// возвращаем встроенные в grpc коды ошибки
		return status.Error(codes.InvalidArgument, "creator id required")
	}

	// также проверяем на валидность uuid
	if _, err := uuid.Parse(req.GetCreatorId()); err != nil {
		// возвращаем встроенные в grpc коды ошибки
		return status.Error(codes.InvalidArgument, "wrong creator id")
	}

	return nil
}
