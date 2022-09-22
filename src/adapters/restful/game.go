package restful

import (
	"hexagonal/src/core/dto"
	"hexagonal/src/core/ports"

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

func (handler *restful) Get(c *gin.Context) {
	game, err := handler.gamePort.Get(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, game)
}

func (handler *restful) Create(c *gin.Context) {
	body := dto.BodyCreate{}
	c.BindJSON(&body)

	game, err := handler.gamePort.Create(body.Name, body.Size, body.Bombs)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, dto.BuildResponseCreate(game))
}

func (handler *restful) RevealCell(c *gin.Context) {
	body := dto.BodyRevealCell{}
	c.BindJSON(&body)

	game, err := handler.gamePort.Reveal(c.Param("id"), body.Row, body.Col)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, dto.BuildResponseRevealCell(game))
}
