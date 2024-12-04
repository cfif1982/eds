package document

import (
	"context"
	"fmt"

	"github.com/cfif1982/eds/internal/models"
	"github.com/google/uuid"
)

func (s *Services) GetById(
	ctx context.Context,
	documentID uuid.UUID,
) (*models.Document, error) {

	// находим документ по его id
	doc, err := s.docRepo.GetByID(ctx, documentID)

	if err != nil {
		return nil, fmt.Errorf("GetById() service error: %w", err)
	}

	return doc, nil
}
