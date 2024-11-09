package bootstraper

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/cfif1982/eds/internal/config"
	"github.com/cfif1982/eds/internal/eds/infrastructure/handlers"
	"github.com/cfif1982/eds/internal/eds/infrastructure/server"
	docRepo "github.com/cfif1982/eds/internal/eds/repositories/document"
	docUseCases "github.com/cfif1982/eds/internal/eds/useCases/document"
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
	// Солздаю подключение к БД
	// DSN для СУБД
	databaseDSN := fmt.Sprintf(dbConnFormat, b.cfg.DB.Host, b.cfg.DB.User, b.cfg.DB.Password, b.cfg.DB.Name)

	db, err := sql.Open("pgx", databaseDSN)
	if err != nil {
		panic("sql open error: " + err.Error())
	}

	// создаем репозиторий для документа
	docRepo, err := docRepo.NewPostgresRepo(b.log, b.cfg, db)
	if err != nil {
		panic("Document Repo error: " + err.Error())
	}

	// TODO: создать репозиторий для юзера

	// создаем UseCases для работы с документами
	docUseCases := docUseCases.NewUseCases(b.log, docRepo)

	// создаем хэндлеры
	handlers := handlers.NewHandlers(b.log, docUseCases)

	// создаем сервер grpc
	grpcSrv := server.NewServer(b.log, b.cfg.GRPC.Port, handlers)

	// запускаем сервер в отдельной горутине,
	// это нужно, для того чтобы слушать сообщения от системы о закрытии приложения - graceful shutdown
	go grpcSrv.MustRun()

	// graceful shutdown
	//**************************************

	// созадем канал, в который будут помещаться сообщения от ОС
	stop := make(chan os.Signal, 1)

	// Функция Notify ждет перечисленные сигналы от ОС и помещает их в указанный канал
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	// ждем прихода данных в канал
	signalType := <-stop

	b.log.Info("stopping application", slog.String("signal", signalType.String()))

	// если данные пришли в канал stop, то нужно завершить приложение
	// вызываем останов сервера
	grpcSrv.Stop()

	b.log.Info("application stoped")

}
