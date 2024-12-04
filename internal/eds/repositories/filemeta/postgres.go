package filemeta

import (
	"database/sql"
	"log/slog"

	"github.com/cfif1982/eds/internal/config"
)

type PostgresRepo struct {
	log        *slog.Logger
	db         *sql.DB
	reqTimeOut int
}

func NewPostgresRepo(log *slog.Logger, cfg *config.Config, db *sql.DB) (*PostgresRepo, error) {
	// TODO: Т.к. не делаю миграцию, то нужно пинговать базу

	return &PostgresRepo{
		log:        log,
		db:         db,
		reqTimeOut: cfg.DB.ReqTimeOut,
	}, nil
}
