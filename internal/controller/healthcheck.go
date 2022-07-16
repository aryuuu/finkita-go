package controller

import (
	"net/http"

	"github.com/aryuuu/finkita/internal/domain"
	"github.com/labstack/echo/v4"
)

// TODO: add db healthcheck
type healthcheckHandler struct{}

func InitHealthCheckHandler(e *echo.Group) {
	handler := &healthcheckHandler{}
	e.GET("/liveness", handler.livenessCheckHandler)
}

func (h *healthcheckHandler) livenessCheckHandler(c echo.Context) error {
	body := domain.Healthcheck{
		Message: "OK",
	}
	return c.JSON(http.StatusOK, body)
}
