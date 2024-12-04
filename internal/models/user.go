package models

import (
	"errors"

	"github.com/google/uuid"
)

// нужно опрделиться , какие ошибки будуо обрабатываться бизнес логикой и выводиться юзеру
var (
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	ID        uuid.UUID
	Email     string
	Name      string
	Telephone string
	OpenKey   string
}
