package document

import (
	"database/sql"
	"fmt"
	"log/slog"
	"path/filepath"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	"github.com/cfif1982/eds/internal/config"
)

const dbConnFormat = "host=%s user=%s password=%s dbname=%s sslmode=disable"

type PostgresRepo struct {
	log        *slog.Logger
	db         *sql.DB
	reqTimeOut int
}

func NewPostgresRepo(log *slog.Logger, cfg *config.Config) (*PostgresRepo, error) {
	// DSN для СУБД
	// Не стал делать общую БД для всех репозиториев, т.к. я могу захотеть для разных репозиториев использолвать разные БД
	// Поэтому саму БД храню в репозитории
	databaseDSN := fmt.Sprintf(dbConnFormat, cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name)

	db, err := sql.Open("pgx", databaseDSN)
	if err != nil {
		return nil, err
	}

	// начинаю миграцию
	// Т.к. делаю миграцию, то не нужно пинговать базу
	log.Info("Start migrating bank database")

	if err = goose.SetDialect("postgres"); err != nil {
		log.Info(err.Error())
	}

	projectRoot, _ := filepath.Abs("../../")
	migrationFolder := filepath.Join(projectRoot, "migrations")

	// // узнаю текущую папку, чтобы передать путь к папке с миграциями
	// ex, err := os.Executable()
	// if err != nil {
	// 	log.Info(err.Error())
	// }
	// exPath := filepath.Dir(ex)

	// exPath += "/migrations"

	err = goose.Up(db, migrationFolder)
	if err != nil {
		log.Info(err.Error() + ": " + migrationFolder)
	}

	log.Info("migrating bank database finished")

	return &PostgresRepo{
		log:        log,
		db:         db,
		reqTimeOut: cfg.DB.ReqTimeOut,
	}, nil
}
