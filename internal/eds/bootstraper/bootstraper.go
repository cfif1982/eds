package bootstraper

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jackc/pgx/v5/stdlib"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/cfif1982/eds/internal/config"
	docHandlers "github.com/cfif1982/eds/internal/eds/infrastructure/handlers/document"
	"github.com/cfif1982/eds/internal/eds/infrastructure/server"
	docRepo "github.com/cfif1982/eds/internal/eds/repositories/document"
	fileRepo "github.com/cfif1982/eds/internal/eds/repositories/file"
	metaRepo "github.com/cfif1982/eds/internal/eds/repositories/filemeta"
	userRepo "github.com/cfif1982/eds/internal/eds/repositories/user"
	docServices "github.com/cfif1982/eds/internal/eds/services/document"
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
	dRepo, err := docRepo.NewPostgresRepo(b.log, b.cfg, db)
	if err != nil {
		panic("Document Repo error: " + err.Error())
	}

	// создаем репозиторий для юзера
	uRepo, err := userRepo.NewPostgresRepo(b.log, b.cfg, db)
	if err != nil {
		panic("User Repo error: " + err.Error())
	}

	// создаем хранилище s3 для файлов
	cfgS3, err := awsConfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("Ошибка загрузки конфигурации для S3: " + err.Error())
	}

	// Создание клиента S3
	s3Client := s3.NewFromConfig(cfgS3)

	// создаем репозиторий для файлов
	fRepo, err := fileRepo.NewS3Repo(b.log, s3Client)
	if err != nil {
		panic("File Repo error: " + err.Error())
	}

	// создаем репозиторий для мета информации файла
	mRepo, err := metaRepo.NewPostgresRepo(b.log, b.cfg, db)
	if err != nil {
		panic("FileMeta Repo error: " + err.Error())
	}

	// создаем Services
	dServices := docServices.NewServices(
		dRepo,
		uRepo,
		mRepo,
		fRepo,
		b.log,
	)

	// создаем хэндлеры
	dHandlers := docHandlers.NewHandlers(b.log, dServices)

	// создаем сервер grpc
	grpcSrv := server.NewServer(b.log, b.cfg.GRPC.Port, dHandlers)

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
