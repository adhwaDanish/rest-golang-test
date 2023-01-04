package main

import (
	"fmt"
	"html/template"
	"io"
	"training-api/db"
	"training-api/routes"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	fmt.Println("running db conn")
	db.Init()
	fmt.Println("running API")
	e := routes.Init()
	//render all template
	renderer := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = renderer
	e.Logger.Fatal(e.Start(":1234"))
}
