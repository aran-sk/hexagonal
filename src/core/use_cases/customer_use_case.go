package use_cases

import (
	"hexagonal/src/config/app_errors"
	"hexagonal/src/core/domain"
	"hexagonal/src/core/ports"
	"hexagonal/src/helpers/uuid"

	"github.com/matiasvarela/errors"
)

type CustomerUseCase struct {
	customersRepository ports.CustomerRepositoryPort
	uuid                uuid.Generator
}

func NewCustomerUseCase(customersRepository ports.CustomerRepositoryPort, uuid uuid.Generator) *CustomerUseCase {
	return &CustomerUseCase{
		customersRepository: customersRepository,
		uuid:                uuid,
	}
}

func (customerUseCase *CustomerUseCase) Get(id string) (domain.Customer, error) {
	customer, err := customerUseCase.customersRepository.Get(id)
	if err != nil {
		if errors.Is(err, app_errors.NotFound) {
			return domain.Customer{}, errors.New(app_errors.NotFound, err, "customer not found")
		}

		return domain.Customer{}, errors.New(app_errors.Internal, err, "get customer from repository has failed")
	}

	return customer, nil
}

func (customerUseCase *CustomerUseCase) Create(name string, surname string) (domain.Customer, error) {
	customer := domain.NewCustomer(customerUseCase.uuid.NewUUID(), name, surname)
	if stataus, err := customerUseCase.customersRepository.Save(customer); err != nil && stataus != "Ok" {
		return domain.Customer{}, errors.New(app_errors.Internal, err, "create customer into repository has failed")
	}

	return customer, nil
}

func (customerUseCase *CustomerUseCase) Update(id string, name string, surname string) (string, error) {
	customer := domain.NewCustomer(id, name, surname)
	if stataus, err := customerUseCase.customersRepository.Save(customer); err != nil && stataus != "OKAY" {
		return "FAILED", errors.New(app_errors.Internal, err, "create customer into repository has failed")
	}

	return "OKAY", nil
}

func (customerUseCase *CustomerUseCase) Delete(id string) (string, error) {
	if stataus, err := customerUseCase.customersRepository.Delete(id); err != nil && stataus != "OKAY" {
		return "FAILED", errors.New(app_errors.Internal, err, "delete customer from repository has failed")
	}

	return "OKAY", nil
}
