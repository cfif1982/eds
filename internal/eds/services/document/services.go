package document

import (
	"context"
	"log/slog"

	"github.com/cfif1982/eds/internal/models"
	"github.com/google/uuid"
)

// Q: вопрос по наименованию функций.
// как тут лучше Save или Update?
// Add, Create, Put
// Get, Read
type DocumentRepo interface {
	Add(ctx context.Context, doc *models.Document) error
	Update(ctx context.Context, doc *models.Document) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Document, error)
}

type UserRepo interface {
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.User, error)
}

// Q: долделать интерфейс
type FileRepo interface {
	Get(ctx context.Context, key string) error            // тут должен вернуть физически файл?
	Put(ctx context.Context, file string) (string, error) // тут передаем физически файл?
}

type FileMetaRepo interface {
	Add(ctx context.Context, file *models.File) error
	Get(ctx context.Context, key string) (*models.File, error)
	Update(ctx context.Context, file *models.File) error
}

type Services struct {
	log          *slog.Logger
	docRepo      DocumentRepo
	userRepo     UserRepo
	fileRepo     FileRepo
	fileMetaRepo FileMetaRepo
}

// Q: тут у меня получается общий сервис. Что-то я запутался в структуре((
func NewServices(
	docRepo DocumentRepo,
	userRepo UserRepo,
	fileMetaRepo FileMetaRepo,
	fileRepo FileRepo,
	log *slog.Logger,
) *Services {
	return &Services{
		log:          log,
		docRepo:      docRepo,
		userRepo:     userRepo,
		fileRepo:     fileRepo,
		fileMetaRepo: fileMetaRepo,
	}
}
