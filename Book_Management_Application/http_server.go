package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func viewMainPage(c echo.Context) error {

	argument := c.Request().URL.Query()

	fmt.Print(argument)

	if len(argument) != 0 {
		title := argument.Get("title")

		addBook(string(title))

		var h template.HTML = template.HTML("")

		for _, book := range BooksList {
			h = h + "<li>" + template.HTML(book.Title) + "</li>"
		}

		return c.Render(http.StatusOK, "hello", h)

	} else {

		GetBooksList()
		var h template.HTML = template.HTML("")
		for _, book := range BooksList {
			h = h + "<li>" + template.HTML(book.Title) + "</li>"
		}

		return c.Render(http.StatusOK, "hello", h)

	}

}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	t := &Template{
		templates: template.Must(template.ParseGlob("index.html")),
	}

	e.Renderer = t

	e.GET("/", viewMainPage)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))

}
