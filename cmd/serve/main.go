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
			gameRepositoryPort := kvs.New()
			gameUseCase := use_cases.New(gameRepositoryPort, uuid.New())
			gameHandler := restful.New(gameUseCase)

			games.GET(":id", gameHandler.Get)
			games.POST("", gameHandler.Create)
			games.PUT(":id", gameHandler.RevealCell)
		}

		customers := api.Group("/customers")
		{
			customerHandler := restful.CustomerHandler{}
			customers.GET(":id", customerHandler.GetCustomer)
			customers.GET("", customerHandler.ListCustomers)
			customers.POST("", customerHandler.CreateCustomer)
			customers.DELETE(":id", customerHandler.DeleteCustomer)
			customers.PATCH(":id", customerHandler.UpdateCustomer)
		}
	}

	log.Fatal((router.Run(":8080")))
}
