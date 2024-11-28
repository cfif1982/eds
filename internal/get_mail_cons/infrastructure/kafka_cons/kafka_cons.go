package kafkacons

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/IBM/sarama"
	"github.com/cfif1982/eds/internal/get_mail_cons/infrastructure/client"
)

const messageReceived = "Получено сообщение: %s, partition: %d, offset: %d\n"

const topicKeyAddDocument string = "add_document" //  ключ в топике

type GetMailConsumer struct {
	consumer          sarama.Consumer
	partitionConsumer sarama.PartitionConsumer
	client            *client.Client
	log               *slog.Logger
}

func NewKafkaConsumer(
	client *client.Client,
	log *slog.Logger,
	topicName,
	host string,
) (*GetMailConsumer, error) {
	const op = "kafka_cons.NewKafkaConsumer"

	// Настройка консюмера
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest // Читаем только новые сообщения

	// Создание консюмера
	consumer, err := sarama.NewConsumer([]string{host}, config)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Подписка на топик. Читаем только новые сообщения
	partitionConsumer, err := consumer.ConsumePartition(topicName, 0, sarama.OffsetNewest)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &GetMailConsumer{
		consumer:          consumer,
		partitionConsumer: partitionConsumer,
		client:            client,
		log:               log,
	}, nil
}

func (k *GetMailConsumer) Close() {
	k.consumer.Close()
	k.partitionConsumer.Close()
}

func (k *GetMailConsumer) Run(ctx context.Context) {
	// Чтение сообщений
	go func() {
		k.log.Info("Ожидание сообщений...")
		for message := range k.partitionConsumer.Messages() {
			mes := fmt.Sprintf(messageReceived, string(message.Value), message.Partition, message.Offset)
			k.log.Info(mes)
			err := k.checkMessage(ctx, message)
			if err != nil {
				k.log.Error("Ошибка при обработке сообщения: %v", err)
			}
		}
	}()
}

// проверяем полученное сообщение.
func (k *GetMailConsumer) checkMessage(ctx context.Context, message *sarama.ConsumerMessage) error {
	key := string(message.Key)

	switch key {
	// получили сообщение от кафки о созданни документа
	case topicKeyAddDocument:
		// это шаблон для других топиков
		//************************************************
		// Десериализация JSON в структуру
		// var msg *models.Document
		// err := json.Unmarshal(message.Value, msg)

		// // тут ллогика другая - нужно получить id создателя документа
		// if err != nil {
		// 	k.log.Info("Ошибка при десериализации сообщения: %v", err)
		// 	return
		// }
		//************************************************

		var msg string

		// Десериализация JSON
		err := json.Unmarshal(message.Value, msg)

		if err != nil {

			return fmt.Errorf("Ошибка при десериализации сообщения: %v", err)
		}

		// делаем grpc запрос в eds
		k.client.AddDocument(ctx, msg)

	}

	return nil
}
