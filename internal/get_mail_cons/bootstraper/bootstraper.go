package bootstraper

import (
	"context"
	"log/slog"

	"github.com/cfif1982/eds/internal/config"
	"github.com/cfif1982/eds/internal/get_mail_cons/infrastructure/client"
	kafkacons "github.com/cfif1982/eds/internal/get_mail_cons/infrastructure/kafka_cons"
)

const kafkaTopicName = "mail_topic"

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

	// Создаем grpc клиента
	getMailClient, err := client.NewClient(
		context.Background(),
		b.log,
		b.cfg.ClientsCfg.GetMail.Address,
		b.cfg.ClientsCfg.GetMail.Timeout,
		b.cfg.ClientsCfg.GetMail.RetriesCount,
	)

	if err == nil {
		b.log.Error("failed to init get mail client", err)
		return
	}

	// Создаем слушателя кафки
	consumer, err := kafkacons.NewKafkaConsumer(
		getMailClient,
		b.log,
		kafkaTopicName,
		b.cfg.KafkaHost,
	)

	if err == nil {
		b.log.Error("failed to init get mail kafka consumer", err)
		return
	}

	defer consumer.Close()

	// запускаем слушателя сообщений от кафки
	// TODO: сделать передачу контекста для закрытия
	ctx := context.TODO()
	consumer.Run(ctx)
}
