package uuid

import (
	"github.com/google/uuid"
)

type Generator interface {
	New() string
}

type uuidGenerator struct{}

func New() Generator {
	return &uuidGenerator{}
}

func (u uuidGenerator) New() string {
	return uuid.New().String()
}
