package restful

import (
	"hexagonal/src/core/dto"
	"hexagonal/src/core/ports"

	"net/http"

	"github.com/gin-gonic/gin"
)

type restful struct {
	gamePort ports.GamePort
}

func New(gameUseCase ports.GamePort) *restful {
	return &restful{
		gamePort: gameUseCase,
	}
}

func (handler *restful) Get(context *gin.Context) {
	game, err := handler.gamePort.Get(context.Param("id"))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, game)
}

func (handler *restful) Create(context *gin.Context) {
	body := dto.BodyCreate{}
	context.BindJSON(&body)

	game, err := handler.gamePort.Create(body.Name, body.Size, body.Bombs)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, dto.BuildResponseCreate(game))
}

func (handler *restful) RevealCell(context *gin.Context) {
	body := dto.BodyRevealCell{}
	context.BindJSON(&body)

	game, err := handler.gamePort.Reveal(context.Param("id"), body.Row, body.Col)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, dto.BuildResponseRevealCell(game))
}
