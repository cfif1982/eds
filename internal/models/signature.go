package models

import (
	"time"

	"github.com/google/uuid"
)

type Signature struct {
	ID            uuid.UUID
	Signer        uuid.UUID
	SignatureFile string
	Date          time.Time
}

func NewSignature(
	id, signer uuid.UUID,
	signatureFile string,
	date time.Time,
) *Signature {

	return &Signature{
		ID:            id,
		Signer:        signer,
		SignatureFile: signatureFile,
		Date:          date,
	}
}
