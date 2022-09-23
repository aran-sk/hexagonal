package restful

import (
	"hexagonal/src/core/domain"
	"hexagonal/src/core/ports"

	"net/http"

	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	customerPort ports.CustomerPort
}

func NewCustomerHandler(customerUseCase ports.CustomerPort) *customerHandler {
	return &customerHandler{
		customerPort: customerUseCase,
	}
}

func (customerHandler *customerHandler) GetCustomer(context *gin.Context) {
	customer, err := customerHandler.customerPort.Get(context.Param("id"))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, customer)
}

func (customerHandler *customerHandler) CreateCustomer(context *gin.Context) {
	body := domain.Customer{}
	context.BindJSON(&body)

	customer, err := customerHandler.customerPort.Create(body.Name, body.Surname)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, customer)
}

func (customerHandler *customerHandler) DeleteCustomer(context *gin.Context) {
	stataus, err := customerHandler.customerPort.Delete(context.Param("id"))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, stataus)
}

func (customerHandler *customerHandler) UpdateCustomer(context *gin.Context) {
	body := domain.Customer{}
	context.BindJSON(&body)

	stataus, err := customerHandler.customerPort.Update(body.ID, body.Name, body.Surname)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, stataus)
}
