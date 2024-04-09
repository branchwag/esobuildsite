package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("*.html")),
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = NewTemplates()

	e.GET("/", func(c echo.Context) error {
		
		return c.Render(http.StatusOK, "index", "data")
	})

	e.GET("/sorcerer", func(c echo.Context) error {
		return c.Render(http.StatusOK, "sorcerer", "buildinfo")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
