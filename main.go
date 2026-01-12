package main

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data any, c echo.Context) error {
 return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("pages/*.html")),
	}
}

type Count struct {
	 Count int
} 

func main() {
	e := echo.New()
	e.Use(middleware.RequestLogger())
	
	e.Renderer = newTemplate()
	
	// count := Count { Count: 0 }
	
	inner := func() {
        fmt.Println("hi")
    }
    inner()
	
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", inner)
	})
	
	e.Logger.Fatal(e.Start(":3000"))
}