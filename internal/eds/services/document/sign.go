package document

import (
	"context"

	"github.com/google/uuid"
)

func (s *Services) SignDocument(
	ctx context.Context,
	documentID uuid.UUID,
	signerEmail string,
	signatureFileURL string,
) error {

}
