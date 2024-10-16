package models

import (
	"time"

	"github.com/google/uuid"
)

type Signature struct {
	id                uuid.UUID
	signer            uuid.UUID
	signatureFileName string
	date              time.Time
}
