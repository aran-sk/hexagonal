package restful

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(context *gin.Context) {
	context.String(http.StatusOK, http.StatusText(http.StatusOK))
}
