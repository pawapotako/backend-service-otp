package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type defaultHandler struct {
}

func NewDefaultHandler(e *echo.Echo) {
	handler := defaultHandler{}

	v1 := e.Group("/v1")
	v1.GET("/otps/health-check", handler.healthCheck)
}

// HealthCheck godoc
// @Summary Check API health
// @Description Check if the API server is running
// @Tags Health Check
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /health-check [get]
func (h defaultHandler) healthCheck(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, map[string]string{"message": "ok"})
}
