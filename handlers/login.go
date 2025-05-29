package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterLoginRoutes(e *echo.Echo) {
	e.GET("/login", loginPage)
	e.POST("/login", loginSubmit)
}

func loginPage(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", nil)
}

func loginSubmit(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "admin" && password == "1234" {
		return c.HTML(http.StatusOK, "<p style='color:green;'>Login efetuado com sucesso!</p>")
	}
	return c.HTML(http.StatusUnauthorized, "<p style='color:red;'>Usuário ou senha inválidos.</p>")
}
