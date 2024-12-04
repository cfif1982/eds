package models

import "github.com/google/uuid"

type File struct {
	ID       uuid.UUID
	FileName string
}

func NewFile(
	id uuid.UUID,
	url string,
) *File {

	return &File{
		ID:       id,
		FileName: url,
	}
}
