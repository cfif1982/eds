package documents

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"

	"github.com/cfif1982/eds/pkg/logger"
	"github.com/pressly/goose/v3"
)

// postgres репозиторий.
type PostgresRepository struct {
	db *sql.DB
}

// Создаем репозиторий БД.
func NewPostgresRepository(ctx context.Context, databaseDSN string, logger *logger.Logger) (*PostgresRepository, error) {

	db, err := sql.Open("pgx", databaseDSN)
	if err != nil {
		return nil, err
	}

	// начинаю миграцию
	// Т.к. делаю миграцию, то не нужно пинговать базу
	logger.Info("Start migrating database")

	if err = goose.SetDialect("postgres"); err != nil {
		logger.Info(err.Error())
	}

	// узнаю текущую папку, чтобы передать путь к папке с миграциями
	ex, err := os.Executable()
	if err != nil {
		logger.Info(err.Error())
	}
	exPath := filepath.Dir(ex)

	exPath += "/migrations"

	err = goose.Up(db, exPath)
	if err != nil {
		logger.Info(err.Error() + ": " + exPath)
	}

	logger.Info("migrating database finished")

	return &PostgresRepository{
		db: db,
	}, nil
}
