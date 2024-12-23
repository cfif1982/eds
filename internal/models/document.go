package models

import (
	"time"

	"github.com/google/uuid"
)

type Document struct {
	id      uuid.UUID
	creator uuid.UUID
	signers []uuid.UUID
	files   []File
	approve bool
	date    time.Time
}

func NewDocument(
	id,
	creator uuid.UUID,
	signers []uuid.UUID,
	files []File,
	approve bool,
	date time.Time,
) *Document {
	return &Document{
		id:      id,
		creator: creator,
		signers: signers,
		files:   files,
		approve: approve,
		date:    date,
	}
}

func CreateDocument(creator uuid.UUID) *Document {
	var signers []uuid.UUID
	uuid := uuid.New()
	files := []File{}

	return NewDocument(
		uuid,
		creator,
		signers,
		files,
		false,
		time.Now(),
	)
}

func (d *Document) ID() uuid.UUID {
	return d.id
}

func (d *Document) Creator() uuid.UUID {
	return d.creator
}

func (d *Document) Signers() []uuid.UUID {
	return d.signers
}

func (d *Document) Files() []File {
	return d.files
}

func (d *Document) Approve() bool {
	return d.approve
}

func (d *Document) Date() time.Time {
	return d.date
}
