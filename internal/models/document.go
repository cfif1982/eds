package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Q: примеры ошибок
// нужно опрделиться , какие ошибки будуо обрабатываться бизнес логикой и выводиться юзеру
var (
	ErrDocumentNotFound     = errors.New("document not found")
	ErrDocumentAlreadyExist = errors.New("document already exist")
)

// Q: ну и тут тоже вопрос - по DDD поле Files нужно хранить как список объектов File. Но можно же хранить и как список id этих файлов
// как лучше то хранить тогда?
// ведь подписантов мы храним как список id, т.к. в ddd это должны были быть агрегаты, а их храним как id
type Document struct {
	ID      uuid.UUID
	Creator uuid.UUID
	Signers []uuid.UUID
	Files   []File
	Approve bool
	Date    time.Time
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
		ID:      id,
		Creator: creator,
		Signers: signers,
		Files:   files,
		Approve: approve,
		Date:    date,
	}
}
