package document

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/cfif1982/eds/internal/models"
)

// добавить документ
func (b *PostgresRepo) Add(ctx context.Context, doc *models.Document) error {
	// настраиваем squirrel для работы с postgres
	psq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	// создаем запрос на добавление документа
	query, args, _ := psq.
		Insert("documents").Columns("id", "creator_id", "date").
		Values(doc.ID(), doc.Creator(), doc.Date()).
		ToSql()

	// создаю контекст для запроса
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(b.reqTimeOut)*time.Second)
	defer cancel()

	// выполняю запрос
	_, err := b.db.ExecContext(ctxTimeout, query, args...)

	// Q: работа с ошибками
	// оборачиываю ошибку и возвращаю наверх в useCase
	if err != nil {
		// у меня в models есть свои ошибки, например например models.ErrUserNotFound
		// здесь мне нужно проверить - полученная ошибка - это ошибка пользователь не найден
		// если да, то возвращаю models.ErrUserNotFound, если нет, то оборачиваю полученную ошибку в текст
		// if err == ошибка юезр не найден{return models.ErrUserNotFound} else
		// либо же здесь это не проверять, а просто вернуть ошибку выше - пусть там разбираются
		// но с другой стороны - это неправильно. Зачем слою выше знать об ошибках репозитория
		// я могу поменять репозиторий и тогда придется там менять логику обработки ошибки
		return fmt.Errorf("failed to add document, repo error: %w", err)
	}

	return nil
}
