package main

import (
	"hexagonal/src/adapters/http"
	"hexagonal/src/adapters/repositories/memory_kvs"
	"hexagonal/src/config/uuid"
	"hexagonal/src/core/use_cases"

	"github.com/gin-gonic/gin"
)

func main() {
	gameRepositoryPort := memory_kvs.NewMemKVS()
	gameUseCase := use_cases.New(gameRepositoryPort, uuid.New())
	gameUsingHttp := http.NewHTTPHandler(gameUseCase)

	router := gin.New()
	router.GET("/games/:id", gameUsingHttp.Get)
	router.POST("/games", gameUsingHttp.Create)
	router.PUT("/games/:id", gameUsingHttp.RevealCell)

	router.Run(":8080")
}
