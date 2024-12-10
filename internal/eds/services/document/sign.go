package document

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/cfif1982/eds/internal/models"
	"github.com/google/uuid"
)

func (s *Services) SignDocument(
	ctx context.Context,
	documentID uuid.UUID,
	signerEmail string,
	signatureFileURL string,
) error {
	// находим юзера по email
	signer, err := s.userRepo.GetByEmail(ctx, signerEmail)
	if err != nil {
		return fmt.Errorf("SignDocument() service error: %w", err)
	}

	// находим документ по его id
	doc, err := s.docRepo.GetByID(ctx, documentID)
	if err != nil {
		return fmt.Errorf("SignDocument() service error: %w", err)
	}

	// проверяем подпись юзера
	err = checkSign(signer, doc, signatureFileURL)
	if err != nil {
		return fmt.Errorf("SignDocument() service error: %w", err)
	}

	// Если всё ок, то создаем сигнатуру и сохраняем документ
	signature := models.NewSignature(uuid.New(), signer.ID, signatureFileURL, time.Now())
	doc.Signatures = append(doc.Signatures, *signature)

	// сохраняем документ в БД
	err = s.docRepo.Update(ctx, doc)

	if err != nil {
		return fmt.Errorf("SignDocument() service error: %w", err)
	}

	s.log.Info("document signed", slog.Any("documentID", documentID))

	// TODO отправляем следуюещму подписанту

	return nil
}

func checkSign(signer *models.User, doc *models.Document, signatureFileURL string) error {
	// TODO доделать проверку подписи

	return nil
}
