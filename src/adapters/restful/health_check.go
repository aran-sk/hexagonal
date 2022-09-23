package restful

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckHandler struct {
}

func (h *HealthCheckHandler) HealthCheck(context *gin.Context) {
	context.String(http.StatusOK, http.StatusText(http.StatusOK))
}
