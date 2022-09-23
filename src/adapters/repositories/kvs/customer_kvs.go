package kvs

import (
	"encoding/json"
	"hexagonal/src/config/app_errors"
	"hexagonal/src/core/domain"

	"github.com/matiasvarela/errors"
)

type CustomerKeyValueStore struct {
	kvs map[string][]byte
}

func NewCustomerKeyValueStore() *CustomerKeyValueStore {
	return &CustomerKeyValueStore{
		kvs: map[string][]byte{},
	}
}

func (repo *CustomerKeyValueStore) Get(id string) (domain.Customer, error) {
	if value, ok := repo.kvs[id]; ok {
		customer := domain.Customer{}
		err := json.Unmarshal(value, &customer)
		if err != nil {
			return domain.Customer{}, errors.New(app_errors.Internal, err, "fail to get value from kvs")
		}

		return customer, nil
	}

	return domain.Customer{}, errors.New(app_errors.NotFound, nil, "customer not found in kvs")
}

func (repo *CustomerKeyValueStore) Save(customer domain.Customer) (string, error) {
	bytes, err := json.Marshal(customer)
	if err != nil {
		return "FAILED", errors.New(app_errors.Internal, err, "customer fails at marshal into json string")
	}

	repo.kvs[customer.ID] = bytes
	return "OKAY", nil
}

func (repo *CustomerKeyValueStore) Delete(id string) (string, error) {
	delete(repo.kvs, id)
	return "OKAY", nil
}
