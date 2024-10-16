package bootstraper

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/cfif1982/eds/internal/config"
	"github.com/cfif1982/eds/internal/controller"
	"github.com/cfif1982/eds/internal/infrastructure/handlers"
	"github.com/cfif1982/eds/internal/infrastructure/server"
	docRepo "github.com/cfif1982/eds/internal/repositories/document"
)

type Bootstraper struct {
	cfg        *config.Config
	log        *slog.Logger
	grpcSrv    *server.Server
	controller *controller.Controller
	handlers   *handlers.Handlers
}

func NewBootstraper(cfg *config.Config, log *slog.Logger) *Bootstraper {
	return &Bootstraper{
		cfg: cfg,
		log: log,
	}
}

func (b *Bootstraper) Run() {
	// создаем репозиторий для документа
	docRepo, err := docRepo.NewPostgresRepo(b.log, b.cfg)
	if err != nil {
		panic("Document Repo error: " + err.Error())
	}

	// TODO: создать репозиторий для юзера

	// создаем контроллер
	b.controller = controller.NewController(b.log, docRepo)

	// создаем хэндлеры
	b.handlers = handlers.NewHandlers(b.log, b.controller)

	// создаем сервер grpc
	b.grpcSrv = server.NewServer(b.log, b.cfg.GRPC.Port, b.handlers)

	// запускаем сервер в отдельной горутине,
	// это нужно, для того чтобы слушать сообщения от системы о закрытии приложения - graceful shutdown
	go b.grpcSrv.MustRun()

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
	b.grpcSrv.Stop()

	b.log.Info("application stoped")

}
