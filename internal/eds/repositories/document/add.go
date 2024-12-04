package document

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"

	sq "github.com/Masterminds/squirrel"

	"github.com/cfif1982/eds/internal/models"
)

// добавить документ
func (r *PostgresRepo) Add(ctx context.Context, doc *models.Document) error {
	// настраиваем squirrel для работы с postgres
	psq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	// создаем запрос на добавление документа
	query, args, _ := psq.
		Insert("documents").Columns("id", "creator_id", "date").
		Values(doc.ID, doc.Creator, doc.Date).
		ToSql()

	// создаю контекст для запроса
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.reqTimeOut)*time.Second)
	defer cancel()

	// выполняю запрос
	_, err := r.db.ExecContext(ctxTimeout, query, args...)

	// Q: правильно обрабатываю ошибку?
	// оборачиываю ошибку и возвращаю наверх в servise
	if err != nil {
		// для проверки ошибок Postgres использую пакет github.com/jackc/pgerrcode.
		// проверяем имеет ли ошибка тип *pgconn.PgError
		if pgErr, ok := err.(*pgconn.PgError); ok {
			// если ошибка: запись уже существует
			if pgErr.Code == pgerrcode.UniqueViolation {
				return models.ErrDocumentAlreadyExists
			}
		}

		return fmt.Errorf("Add() document repo error: %w", err)
	}

	return nil
}
