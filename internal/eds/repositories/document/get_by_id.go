package document

import (
	"context"

	"github.com/google/uuid"

	"github.com/cfif1982/eds/internal/models"
)

// добавить документ
func (r *PostgresRepo) GetByID(ctx context.Context, id uuid.UUID) (*models.Document, error) {
	// TODO доделать

	return nil, nil
}
