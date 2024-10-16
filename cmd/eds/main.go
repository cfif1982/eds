package main

import (
	"log/slog"
	"os"

	"github.com/cfif1982/eds/bootstraper"
	"github.com/cfif1982/eds/internal/config"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	// загружаем настройки.
	// Если они не загрузятся, то будет паника, т.к. дальше нет смысла работать
	// поэтому вызываем метод MustLoad()
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting application", slog.Any("config", cfg))

	// создаем bootsraper
	bs := bootstraper.NewBootstraper(cfg, log)

	bs.Run() // запуск bootstraper
}

// настройка логгера
func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

// генерирую grpc
// protoc -I protos\proto protos\proto\eds.proto --go_out=protos\gen --go_opt=paths=source_relative --go-grpc_out=protos\gen --go-grpc_opt=paths=source_relative
