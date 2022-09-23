package restful

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheckHandler godoc
// @summary Health Check
// @description Health checking for the service
// @id HealthCheckHandler
// @produce plain
// @response 200 {string} string "OK"
// @router /healthcheck [get]
type HealthCheckHandler struct {
}

func (h *HealthCheckHandler) HealthCheck(context *gin.Context) {
	context.String(http.StatusOK, http.StatusText(http.StatusOK))
}
