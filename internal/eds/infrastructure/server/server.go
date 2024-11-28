package server

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/cfif1982/eds/internal/eds/infrastructure/handlers"
	edsv1 "github.com/cfif1982/eds/protos/gen"
	"google.golang.org/grpc"
)

type Server struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

// создаем новый grpc сервер
func NewServer(
	log *slog.Logger,
	port int,
	handlers *handlers.Handlers,
) *Server {
	// создаем grpc сервер
	gRPCServer := grpc.NewServer()

	// подключаем наши хэндлеры
	edsv1.RegisterEDSServer(gRPCServer, handlers)

	return &Server{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

// запуск сервера. Если сервер не запустился, то паникуем
// Must обозначает обязательность выполнения функции.
// Если она не запустится, то дальше нет смысла работать
// Поэтому метод Run() мы оборачиваем в MustRun()
func (s *Server) MustRun() {
	if err := s.Run(); err != nil {
		panic(err)
	}
}

// запуск сервера
func (s *Server) Run() error {
	// grpc работает на низком уровне - TCP.
	// Поэтому для его работы нужен слушатель порта для этого протокола
	// для этого создаем слушатель порта
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))

	if err != nil {
		return fmt.Errorf("listener init error: %w", err)
	}

	s.log.Info("grpc server is running", slog.String("addr", listener.Addr().String()))

	// запускаем сервер
	// т.е. указываем серверу обрабатывать запросы, которые приходят на указанный в слушателе адрес
	if err = s.gRPCServer.Serve(listener); err != nil {
		return fmt.Errorf("gRPC server start error: %w", err)
	}

	return nil
}

// остановка сервера
func (s *Server) Stop() {
	s.log.Info("stopping grpc server", slog.Int("port", s.port))

	s.gRPCServer.GracefulStop()
}
