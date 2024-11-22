package document

import (
	"context"
	"log/slog"

	"github.com/cfif1982/eds/internal/models"
)

type Repo interface {
	Add(ctx context.Context, bank *models.Document) error
}

type UseCases struct {
	log  *slog.Logger
	repo Repo
}

func NewUseCases(log *slog.Logger, repo Repo) *UseCases {
	return &UseCases{
		log:  log,
		repo: repo,
	}
}
