package internal

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cfif1982/eds/pkg/logger"
	"github.com/go-chi/chi/v5"

	documentsHandler "github.com/cfif1982/eds/internal/application/documents/handlers"
	usersHandler "github.com/cfif1982/eds/internal/application/users/handlers"

	documentsInfra "github.com/cfif1982/eds/internal/infrastructure/documents"
	usersInfra "github.com/cfif1982/eds/internal/infrastructure/users"
)

const databaseConnectTimeout = 5 // таймаут подключения к БД
const databasePassword = "123"
const databaseHost = "localhost"

// структура сервера.
type Server struct {
	logger *logger.Logger
}

// Конструктор Server.
func NewServer(logger *logger.Logger) *Server {
	return &Server{
		logger: logger,
	}
}

// запуск сервера.
func (s *Server) Run(serverAddr string) error {
	// DSN для СУБД
	databaseDSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", databaseHost, `postgres`, databasePassword, `eds`)

	// создаю контекст для подключения БД
	ctx, cancel := context.WithTimeout(context.Background(), databaseConnectTimeout*time.Second)
	defer cancel()

	// Создаем репозиторий для работы с документами
	documentRepo, err := documentsInfra.NewPostgresRepository(ctx, databaseDSN, s.logger)

	if err != nil {
		s.logger.Fatal("can't initialize postgres for documents DB: " + err.Error())
	} else {
		s.logger.Info("postgres for documents DB initialized")
	}

	// Создаем репозиторий для работы с юзерами
	userRepo, err := usersInfra.NewPostgresRepository(ctx, databaseDSN, s.logger)

	if err != nil {
		s.logger.Fatal("can't initialize postgres for users DB: " + err.Error())
	} else {
		s.logger.Info("postgres for users DB initialized")
	}

	// создаем хэндлер документа
	documentHandler := documentsHandler.NewHandler(documentRepo, s.logger)

	// создаем хэндлер юзера
	userHandler := usersHandler.NewHandler(userRepo, s.logger)

	// создаем роутер
	routerChi := chi.NewRouter()

	// назначаем middleware
	// routerChi.Use(middlewares.GzipCompressMiddleware)
	// routerChi.Use(middlewares.GzipDecompressMiddleware)

	// назначаем ручки для документа
	s.setDocumentHandlers(routerChi, documentHandler)

	// назначаем ручки для юзера
	s.setUserHandlers(routerChi, userHandler)

	s.logger.Info("Starting server", "addr", serverAddr)

	// запуск сервера на нужном адресе и с нужным роутером
	return http.ListenAndServe(serverAddr, routerChi)
}

// назначаем ручки для документа.
func (s *Server) setDocumentHandlers(router *chi.Mux, handler *documentsHandler.Handler) {

	// router.Post(`/api/document/add`, handler.AddDocument())
}

// назначаем ручки для юзера.
func (s *Server) setUserHandlers(router *chi.Mux, handler *usersHandler.Handler) {

	// router.Post(`/api/user/add`, handler.AddUser())
}
