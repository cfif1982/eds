package document

import (
	"github.com/cfif1982/eds/internal/models"
	"github.com/google/uuid"
)

func (u *UseCases) Add(creatorID uuid.UUID) (string, error) {
	// создаем документ
	doc := models.CreateDocument(creatorID)

	// сохраняем документ в БД
	err := u.repo.Add(doc)

	if err != nil {
		return "", err
	}

	// генерируем qr-код
	// TODO: сделать генерацию qr кода
	qrCode := "qrcode"

	return qrCode, nil
}
