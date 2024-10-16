package document

import (
	"log/slog"

	"github.com/cfif1982/eds/internal/models"
)

// Q: приходится объявлять этот интерфейс в двух местах: controller и useCases
// Это правильно?
type Repo interface {
	Add(bank *models.Document) error
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
