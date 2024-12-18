package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}