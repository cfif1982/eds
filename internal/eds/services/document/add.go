package document

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/cfif1982/eds/internal/models"
	"github.com/google/uuid"
)

func (s *Services) AddDocument(ctx context.Context, creatorEmail string) (string, error) {
	// находим id юзера по email
	user, err := s.userRepo.GetByEmail(ctx, creatorEmail)

	if err != nil {
		return "", fmt.Errorf("AddDocument() service error: %w", err)
	}

	// создаем документ
	var signers []uuid.UUID
	var files []uuid.UUID
	uuid := uuid.New()

	doc := models.NewDocument(
		uuid,
		user.ID,
		signers,
		files,
		false,
		time.Now(),
	)

	// сохраняем документ в БД
	err = s.docRepo.Add(ctx, doc)

	if err != nil {
		return "", fmt.Errorf("AddDocument() service error: %w", err)
	}

	// генерируем qr-код
	// TODO: сделать генерацию qr кода
	qrCode := "qrcode"

	s.log.Info("qrcode created", slog.Any("creatorID", user.ID))

	return qrCode, nil
}
