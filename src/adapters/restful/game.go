package restful

import (
	"hexagonal/src/core/dto"
	"hexagonal/src/core/ports"

	"net/http"

	"github.com/gin-gonic/gin"
)

type gameHandler struct {
	gamePort ports.GamePort
}

func NewGameHandler(gameUseCase ports.GamePort) *gameHandler {
	return &gameHandler{
		gamePort: gameUseCase,
	}
}

func (handler *gameHandler) Get(context *gin.Context) {
	game, err := handler.gamePort.Get(context.Param("id"))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, game)
}

func (handler *gameHandler) Create(context *gin.Context) {
	body := dto.BodyGameCreate{}
	context.BindJSON(&body)

	game, err := handler.gamePort.Create(body.Name, body.Size, body.Bombs)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, dto.BuildResponseGameCreate(game))
}

func (handler *gameHandler) RevealCell(context *gin.Context) {
	body := dto.BodyGameRevealCell{}
	context.BindJSON(&body)

	game, err := handler.gamePort.Reveal(context.Param("id"), body.Row, body.Col)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, dto.BuildResponseGameRevealCell(game))
}
