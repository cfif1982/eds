package main

import (
	"log/slog"

	"github.com/cfif1982/eds/internal/config"
	migratorbs "github.com/cfif1982/eds/internal/migrator/bootstraper"
	logger "github.com/cfif1982/eds/internal/pkg/logger/slog"
)

func main() {
	// загружаем настройки.
	// Если они не загрузятся, то будет паника, т.к. дальше нет смысла работать
	// поэтому вызываем метод MustLoad()
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)

	log.Info("starting migrator", slog.Any("config", cfg))

	// создаем bootsraper
	bs := migratorbs.NewBootstraper(cfg, log)

	bs.Run() // запуск bootstraper
}
