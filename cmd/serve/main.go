package main

import (
	"hexagonal/src/adapters/health_check"
	"hexagonal/src/adapters/http"
	"hexagonal/src/adapters/repositories/memory_kvs"
	"hexagonal/src/config/uuid"
	"hexagonal/src/core/use_cases"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/healthcheck", health_check.HealthCheckHandler)

	api := router.Group("/api")
	{
		games := api.Group("/games")
		{
			gameRepositoryPort := memory_kvs.NewMemKVS()
			gameUseCase := use_cases.New(gameRepositoryPort, uuid.New())
			gameUsingHttp := http.NewHTTPHandler(gameUseCase)

			games.GET(":id", gameUsingHttp.Get)
			games.POST("", gameUsingHttp.Create)
			games.PUT(":id", gameUsingHttp.RevealCell)
		}
	}

	log.Fatal((router.Run(":8080")))
}
