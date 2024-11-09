package bootstraper

import (
	"log/slog"

	"database/sql"
	"fmt"

	"github.com/cfif1982/eds/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

const dbConnFormat = "host=%s user=%s password=%s dbname=%s sslmode=disable"

type Bootstraper struct {
	cfg *config.Config
	log *slog.Logger
}

func NewBootstraper(cfg *config.Config, log *slog.Logger) *Bootstraper {
	return &Bootstraper{
		cfg: cfg,
		log: log,
	}
}

func (b *Bootstraper) Run() {
	// DSN для СУБД
	databaseDSN := fmt.Sprintf(dbConnFormat, b.cfg.DB.Host, b.cfg.DB.User, b.cfg.DB.Password, b.cfg.DB.Name)

	db, err := sql.Open("pgx", databaseDSN)
	if err != nil {
		panic("migration error: " + err.Error())
	}

	// начинаю миграцию
	b.log.Info("Start migrating database")

	if err = goose.SetDialect("postgres"); err != nil {
		b.log.Info(err.Error())
	}

	err = goose.Up(db, b.cfg.MigrationFolder)
	if err != nil {
		b.log.Info(err.Error() + ": " + b.cfg.MigrationFolder)
	}

	b.log.Info("migrating database finished")
}
