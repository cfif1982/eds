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
