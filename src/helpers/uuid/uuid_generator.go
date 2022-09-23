package uuid

import (
	"github.com/google/uuid"
)

type Generator interface {
	NewUUID() string
}

type uuidGenerator struct{}

func NewUUID() Generator {
	return &uuidGenerator{}
}

func (u uuidGenerator) NewUUID() string {
	return uuid.New().String()
}
