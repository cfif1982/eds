package document

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/cfif1982/eds/internal/models"
	"github.com/google/uuid"
)

func (s *Services) SendDocument(
	ctx context.Context,
	documentID uuid.UUID,
	signersEmail []string,
	filePaths []string,
) error {

	// TODO удаляем все подписи этого документа из БД
	// TODO удаляем все файлы этого документа из БД

	// т.к. файлы физически уже загружены в s3 или локальное хранилище, то сюда нам переданы только слайс ссылок на эти файлы в s3
	// тогда нужно пробежаться по этому слайсу и создать объекты файлов
	files := make([]models.File, 0, len(filePaths))

	for _, url := range filePaths {
		uuid := uuid.New()

		file := models.NewFile(uuid, url)

		files = append(files, *file)
	}

	// получаю слайс id подписантов по их email
	signersID := make([]uuid.UUID, 0, len(signersEmail))

	for _, email := range signersEmail {
		user, err := s.userRepo.GetByEmail(ctx, email)

		if err != nil {
			return fmt.Errorf("SendDocument() service error: %w", err)
		}

		signersID = append(signersID, user.ID)
	}

	// находим документ по его id
	doc, err := s.docRepo.GetByID(ctx, documentID)

	if err != nil {
		return fmt.Errorf("SendDocument() service error: %w", err)
	}

	// меняем параметры документа
	doc.Files = files
	doc.SignersID = signersID

	// сохраняем документ в БД
	err = s.docRepo.Update(ctx, doc)

	if err != nil {
		return fmt.Errorf("SendDocument() service error: %w", err)
	}

	s.log.Info("document updated", slog.Any("documentID", documentID))

	// TODO отправляем сообщение первому подписанту
	// Q: вопрос сообщения подписантам нужно отправлять в заданном порядке
	// я это реализовал следующим образом:
	// чтобы не создавать дополнителную таблицу с подписантами документа и их порядком,
	// я воспользьвался таблицей signatures. Теперь при обновлении данных документа,
	// в таблицу files добавляю файлы, и сразу добавляю записи в таблицу signatures всех подписантов данного файла.
	// но с пустым полем signature_file_name, т.к. файл еще не подписан данным подписантом
	// нормальное решение?
	return nil
}
