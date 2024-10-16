package user

import "github.com/google/uuid"

type User struct {
	id        uuid.UUID
	email     string
	name      string
	telephone string
	openKey   string
}
