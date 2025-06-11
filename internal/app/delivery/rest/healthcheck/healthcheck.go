package healthcheck

import (
	"go-grpc-boilerplate/internal/app/usecase/healthcheck"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      Health Check
// @Description  Get the health status of the service
// @Tags         healthcheck
// @Produce      json
// @Success      200  {object}  healthcheck.InportResponse  "OK response"
// @Failure      503  {object}  healthcheck.InportResponse  "Service Unavailable response"
// @Router       /healthcheck [get]
func HealthCheck(inport healthcheck.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := inport.Execute(c.Copy().Request.Context())

		if resp.Message == "OK" {
			c.JSON(http.StatusOK, resp)
		} else {
			c.JSON(http.StatusServiceUnavailable, resp)
		}
	}
}
