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
	Creator    uuid.UUID
	Signers    []uuid.UUID
	Files      []uuid.UUID
	Signatures []uuid.UUID
	Approve    bool
	Date       time.Time
}

func NewDocument(
	id,
	creator uuid.UUID,
	signers []uuid.UUID,
	files []uuid.UUID,
	approve bool,
	date time.Time,
) *Document {
	return &Document{
		ID:      id,
		Creator: creator,
		Signers: signers,
		Files:   files,
		Approve: approve,
		Date:    date,
	}
}
