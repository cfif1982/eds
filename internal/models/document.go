package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// нужно опрделиться , какие ошибки будуо обрабатываться бизнес логикой и выводиться юзеру
var (
	ErrDocumentNotFound      = errors.New("document not found")
	ErrDocumentAlreadyExists = errors.New("document already exist") // на тот случай, если uuid уже существует
)

type Document struct {
	ID         uuid.UUID
	CreatorID  uuid.UUID
	SignersID  []uuid.UUID
	Files      []File
	Signatures []Signature
	Approve    bool
	Date       time.Time
}

func NewDocument(
	id,
	creator uuid.UUID,
	signers []uuid.UUID,
	files []File,
	signatures []Signature,
	approve bool,
	date time.Time,
) *Document {
	return &Document{
		ID:         id,
		CreatorID:  creator,
		SignersID:  signers,
		Files:      files,
		Signatures: signatures,
		Approve:    approve,
		Date:       date,
	}
}
