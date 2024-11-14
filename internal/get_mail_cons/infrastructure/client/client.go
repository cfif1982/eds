package client

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	edsv1 "github.com/cfif1982/eds/protos/gen"
	grpclog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	gRPCClient edsv1.EDSClient
}

func NewClient(
	ctx context.Context,
	log *slog.Logger,
	addr string,
	timeout time.Duration,
	retriesCount int,
) (*Client, error) {
	const op = "client.NewClient"

	// интерсептер для ретраев
	retryOpts := []grpcretry.CallOption{
		// указываем какие коды нам нужно ретраить
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(retriesCount)),  // количество ретраев
		grpcretry.WithPerRetryTimeout(timeout), // таймаут ретраев
	}

	// интерсептер для логирования
	logOpts := []grpclog.Option{
		// хотим залогировать payload, который получаем и оптравляем
		// т.е. тело запроса и ответа
		grpclog.WithLogOnEvents(grpclog.PayloadReceived, grpclog.PayloadSent),
	}

	// Пытаемся установить соединение с сервером, используя DialContext.
	con, err := grpc.DialContext(ctx, addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()), // про шифрование - без него
		grpc.WithChainUnaryInterceptor( // делаем цепочку из интерсепторов
			grpclog.UnaryClientInterceptor(InterseptorLogger(log), logOpts...),
			grpcretry.UnaryClientInterceptor(retryOpts...),
		),
	)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Client{
		gRPCClient: edsv1.NewEDSClient(con),
	}, nil
}

func InterseptorLogger(l *slog.Logger) grpclog.Logger {
	return grpclog.LoggerFunc(func(ctx context.Context, lvl grpclog.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

func (c *Client) AddDocument(ctx context.Context, creatorID string) error {
	const op = "grpc.AddDocument"

	// нам здесь не нужен результат запроса
	// важно понять, что запрос отработал без ошибок
	// результат запроса будет отправлен в кафку в SendMessageQueue
	_, err := c.gRPCClient.AddNewDocument(ctx, &edsv1.AddNewDocumentRequest{
		CreatorId: creatorID,
	})

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
