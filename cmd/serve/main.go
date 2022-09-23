package main

import (
	"hexagonal/src/adapters/repositories/kvs"
	"hexagonal/src/adapters/restful"
	"hexagonal/src/core/use_cases"
	"hexagonal/src/helpers/uuid"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/healthcheck", restful.HealthCheckHandler)

	api := router.Group("/api")
	{
		games := api.Group("/games")
		{
			gameRepositoryPort := kvs.New()
			gameUseCase := use_cases.New(gameRepositoryPort, uuid.New())
			gameUsingHttp := restful.New(gameUseCase)

			games.GET(":id", gameUsingHttp.Get)
			games.POST("", gameUsingHttp.Create)
			games.PUT(":id", gameUsingHttp.RevealCell)
		}
	}

	log.Fatal((router.Run(":8080")))
}
