package file

import (
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Repo struct {
	log    *slog.Logger
	client *s3.Client
}

func NewS3Repo(log *slog.Logger, client *s3.Client) (*S3Repo, error) {

	return &S3Repo{
		log:    log,
		client: client,
	}, nil
}
