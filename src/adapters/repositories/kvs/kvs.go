package kvs

import (
	"encoding/json"
	"hexagonal/src/config/app_errors"
	"hexagonal/src/config/messages"
	"hexagonal/src/core/domain"

	"github.com/matiasvarela/errors"
)

type KeyValueStore struct {
	kvs map[string][]byte
}

func New() *KeyValueStore {
	return &KeyValueStore{
		kvs: map[string][]byte{},
	}
}

func (repo *KeyValueStore) Get(id string) (domain.Game, error) {

	if value, ok := repo.kvs[id]; ok {
		game := domain.Game{}
		err := json.Unmarshal(value, &game)
		if err != nil {
			return domain.Game{}, errors.New(app_errors.Internal, err, messages.GameNotFoundFromKVS)
		}

		return game, nil
	}

	return domain.Game{}, errors.New(app_errors.NotFound, nil, "game not found in kvs")
}

func (repo *KeyValueStore) Save(game domain.Game) error {
	bytes, err := json.Marshal(game)
	if err != nil {
		return errors.New(app_errors.InvalidInput, err, messages.GameMarshalingFailed)
	}

	repo.kvs[game.ID] = bytes
	return nil
}
