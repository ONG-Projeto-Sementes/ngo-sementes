package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterHomeRoutes(e *echo.Echo) {
	e.GET("/", homePage)
}

func homePage(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", nil)
}
