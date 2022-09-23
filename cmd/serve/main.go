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

	healthCheckHandler := restful.HealthCheckHandler{}
	router.GET("/healthcheck", healthCheckHandler.HealthCheck)

	api := router.Group("/api")
	{
		games := api.Group("/games")
		{
			gameRepositoryPort := kvs.NewGameKeyValueStore()
			gameUseCase := use_cases.NewGameUseCase(gameRepositoryPort, uuid.NewUUID())
			gameHandler := restful.NewGameHandler(gameUseCase)

			games.GET(":id", gameHandler.Get)
			games.POST("", gameHandler.Create)
			games.PUT(":id", gameHandler.RevealCell)
		}

		customers := api.Group("/customers")
		{
			customerRepositoryPort := kvs.NewCustomerKeyValueStore()
			customerUseCase := use_cases.NewCustomerUseCase(customerRepositoryPort, uuid.NewUUID())
			customerHandler := restful.NewCustomerHandler(customerUseCase)
			customers.GET(":id", customerHandler.GetCustomer)
			customers.POST("", customerHandler.CreateCustomer)
			customers.DELETE(":id", customerHandler.DeleteCustomer)
			customers.PATCH(":id", customerHandler.UpdateCustomer)
		}
	}

	log.Fatal((router.Run(":8080")))
}
