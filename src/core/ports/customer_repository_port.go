package ports

import "hexagonal/src/core/domain"

type CustomerRepositoryPort interface {
	Get(id string) (domain.Customer, error)
	Save(customer domain.Customer) (string, error)
	Delete(id string) (string, error)
}
