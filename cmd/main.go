package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"mywebsite.tv/sementes/handlers"
)

// TemplateRenderer para renderizar os templates HTML
type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = renderer

	e.Static("/css", "css")
	e.Static("/images", "images")
	e.Static("/svg", "svg")
	e.Static("/js", "js")

	handlers.RegisterLoginRoutes(e)
	handlers.RegisterHomeRoutes(e)

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	e.Logger.Fatal(e.Start(":42069"))
}
