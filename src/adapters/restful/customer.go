package restful

import (
	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
}

func (h *CustomerHandler) GetCustomer(c *gin.Context) {
	//Logic goes here
}

func (h *CustomerHandler) ListCustomers(c *gin.Context) {
	//Logic goes here
}

func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	//Logic goes here
}

func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
	//Logic goes here
}

func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
	//Logic goes here
}