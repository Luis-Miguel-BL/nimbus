package util

import (
	"github.com/google/uuid"
)

func NewUUID() (id string) {
	id = uuid.New().String()
	return id
}
