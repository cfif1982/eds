package response

import (
	"github.com/google/uuid"
)

type Document struct {
	ID       uuid.UUID
	Creator  User
	Signers  []User
	FileURLs []string
}
