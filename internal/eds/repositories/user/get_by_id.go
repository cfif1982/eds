package user

import (
	"context"

	"github.com/cfif1982/eds/internal/models"
	"github.com/google/uuid"
)

func (r *PostgresRepo) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	// TODO доделать
	return nil, nil
}
