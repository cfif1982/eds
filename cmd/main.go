package main

import (
	// TODO: добавить логгер

	"log"

	"github.com/cfif1982/eds/internal"
	"github.com/cfif1982/eds/pkg/logger"
)

const serverPort = "8080"
const serverHost = "localhost"

func main() {
	// инициализируем логгер
	logger, err := logger.GetLogger()

	// Если логгер не инициализировался, то выводим сообщение с помощью обычного log
	if err != nil {
		log.Fatal("logger zap initialization: FAILURE")
	}

	// выводим сообщенеи об успешной инициализации логгера
	logger.Info("logger zap initialization: SUCCESS")

	// создаем сервер
	srv := internal.NewServer(logger)

	// запускаем сервер
	if err = srv.Run(serverHost + ":" + serverPort); err != nil {
		logger.Fatal("error occured while running http server: " + err.Error())
	}
}
