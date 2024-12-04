package user

import (
	"context"

	"github.com/cfif1982/eds/internal/models"
)

func (r *PostgresRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	// TODO доделать
	return nil, nil
}
