package use_cases

import (
	"hexagonal/src/config/app_errors"
	"hexagonal/src/config/messages"
	"hexagonal/src/core/domain"
	"hexagonal/src/core/ports"
	"hexagonal/src/helpers/uuid"

	"github.com/matiasvarela/errors"
)

type GameUseCase struct {
	gamesRepository ports.GameRepositoryPort
	uuid            uuid.Generator
}

func NewGameUseCase(gamesRepository ports.GameRepositoryPort, uuid uuid.Generator) *GameUseCase {
	return &GameUseCase{
		gamesRepository: gamesRepository,
		uuid:            uuid,
	}
}

func (gameUseCase *GameUseCase) Get(id string) (domain.Game, error) {
	game, err := gameUseCase.gamesRepository.Get(id)
	if err != nil {
		if errors.Is(err, app_errors.NotFound) {
			return domain.Game{}, errors.New(app_errors.NotFound, err, messages.GameNotFound)
		}

		return domain.Game{}, errors.New(app_errors.Internal, err, messages.GameFailedFromRepository)
	}

	game.Board = game.Board.HideBombs()

	return game, nil
}

func (gameUseCase *GameUseCase) Create(name string, size uint, bombs uint) (domain.Game, error) {
	if bombs >= size*size {
		return domain.Game{}, errors.New(app_errors.InvalidInput, nil, messages.GameBombsTooHigh)
	}

	game := domain.NewGame(gameUseCase.uuid.NewUUID(), name, size, bombs)

	if err := gameUseCase.gamesRepository.Save(game); err != nil {
		return domain.Game{}, errors.New(app_errors.Internal, err, messages.GameCannotBeCreatedFromRepository)
	}

	game.Board = game.Board.HideBombs()

	return game, nil
}

func (gameUseCase *GameUseCase) Reveal(id string, row uint, col uint) (domain.Game, error) {

	game, err := gameUseCase.gamesRepository.Get(id)
	if err != nil {
		if errors.Is(err, app_errors.NotFound) {
			return domain.Game{}, errors.New(app_errors.NotFound, err, messages.GameNotFound)
		}

		return domain.Game{}, errors.New(app_errors.Internal, err, messages.GameFailedFromRepository)
	}

	if !game.Board.IsValidPosition(row, col) {
		return domain.Game{}, errors.New(app_errors.InvalidInput, nil, messages.GameInvalidPosition)
	}

	if game.IsOver() {
		return domain.Game{}, errors.New(app_errors.IllegalOperation, nil, messages.GameOver)
	}

	if game.Board.Contains(row, col, domain.CellBomb) {
		game.State = domain.GameStateLost
	} else {
		game.Board.Set(row, col, domain.CellRevealed)

		if !game.Board.IsCellEmpty() {
			game.State = domain.GameStateWon
		}
	}

	if err := gameUseCase.gamesRepository.Save(game); err != nil {
		return domain.Game{}, errors.New(app_errors.Internal, err, messages.GameCannotBeUpdateFromRepository)
	}

	game.Board = game.Board.HideBombs()

	return game, nil
}
