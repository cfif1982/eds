package document

import (
	"context"
	"errors"
	"github.com/cfif1982/eds/internal/models"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

// добавить документ
func (b *PostgresRepo) Add(doc *models.Document) error {
	// настраиваем squirrel для работы с postgres
	psq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	// создаем запрос на добавление документа
	// Q: нужно ли здесь указывать все поля из модели?
	//query := `INSERT INTO documents ("id", "creator_id", "date") VALUES (:id, :creator_id, :date)`

	// sqlx helps us to fill  from models.Document
	// sqlc

	query, args, _ := psq.
		Insert("documents").Columns("id", "creator_id", "date").
		Values(doc.ID(), doc.Creator(), doc.Date()).
		ToSql()

	// создаю контекст для запроса
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(b.reqTimeOut)*time.Second)
	defer cancel()

	// выполняю запрос
	_, err := b.db.ExecContext(ctx, query, args...)
	if err != nil {
		// проверяем ошибку на предмет вставки маршрута с названием, которое уже есть в БД
		// создаем объект *pgconn.PgError - в нем будет храниться код ошибки из БД
		var pgErr *pgconn.PgError

		// преобразуем ошибку к типу pgconn.PgError
		if errors.As(err, &pgErr) {
			// если ошибка- запись существует, то возвращаем эту ошибку
			if pgErr.Code == pgerrcode.UniqueViolation {
				return pgErr
			} else {
				return pgErr
			}
		} else {
			return err
		}
	}

	return nil
}
