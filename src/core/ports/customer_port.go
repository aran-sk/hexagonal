package ports

import "hexagonal/src/core/domain"

type CustomerPort interface {
	Get(id string) (domain.Customer, error)
	Create(name string, surname string) (domain.Customer, error)
	Update(id string, name string, surname string) (string, error)
	Delete(id string) (string, error)
}
