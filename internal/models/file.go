package models

import "github.com/google/uuid"

type File struct {
	id         uuid.UUID
	fileName   string
	signatures []Signature
}
