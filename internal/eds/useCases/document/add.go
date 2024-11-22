package document

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/cfif1982/eds/internal/models"
	"github.com/google/uuid"
)

func (u *UseCases) Add(ctx context.Context, creatorID uuid.UUID) (string, error) {
	// создаем документ
	doc := models.CreateDocument(creatorID)

	// сохраняем документ в БД
	err := u.repo.Add(ctx, doc)

	// Q: работа с ошибками
	if err != nil {
		// здесь у меня идет бизнес логика обработки ошибки
		// я проверяю - если это моя ошибка из models, например models.ErrUserNotFound
		// то я просто верну ошибку дальше по стеку
		// если это ни одна из моих ошибок, то я просто добавлю текст - ошибка в Add useCase и верну дальше
		// или тут в любом случае стоит обернуть ошибку с сообщением о месте ошибки - useCase add document?
		return "", fmt.Errorf("add document useCase error: %w", err)
	}

	// генерируем qr-код
	// TODO: сделать генерацию qr кода
	qrCode := "qrcode"

	// Q: логирую результат в use case. Правильно делаю, что здесь логирую?
	// т.е. в слое репоитория ничегоне логирую, только ошибки передаю
	// а здесь логирую результат
	u.log.Info("qrcode created", slog.Any("creatorID", creatorID))

	return qrCode, nil
}
