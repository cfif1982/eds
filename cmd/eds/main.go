package main

import (
	"log/slog"

	"github.com/cfif1982/eds/internal/config"
	"github.com/cfif1982/eds/internal/eds/bootstraper"
	migratorbs "github.com/cfif1982/eds/internal/migrator/bootstraper"
	logger "github.com/cfif1982/eds/internal/pkg/logger/slog"
)

func main() {
	// загружаем настройки.
	// Если они не загрузятся, то будет паника, т.к. дальше нет смысла работать
	// поэтому вызываем метод MustLoad()
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)

	// запуск мигратора
	mgbs := migratorbs.NewBootstraper(cfg, log) // создаем bootsraper мигратора
	mgbs.Run()                                  // запуск bootstraper мигратора

	log.Info("starting eds", slog.Any("config", cfg))

	// создаем bootsraper приложения
	bs := bootstraper.NewBootstraper(cfg, log)

	// запуск bootstraper приложения
	bs.Run()
}

// генерирую grpc
// protoc -I protos\proto protos\proto\eds.proto --go_out=protos\gen --go_opt=paths=source_relative --go-grpc_out=protos\gen --go-grpc_opt=paths=source_relative
