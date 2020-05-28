package main

import (
	"fmt"
	"github.com/jtarchie/cl-search/cmd/server/views"
	"github.com/jtarchie/cl-search/pkg/load"
	"github.com/jtarchie/cl-search/pkg/parser"
	query "github.com/jtarchie/cl-search/pkg/query"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"
)

//go:generate go run github.com/valyala/quicktemplate/qtc

func main() {
	cities, err := load.AllCities()
	if err != nil {
		log.Fatalf("cannot load YAML: %s", err)
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.Logger())

	e.GET("/", func(context echo.Context) error {
		params, err := parser.ParseParams(context.QueryParam("query"))
		if err != nil {
			return err
		}

		return context.HTML(http.StatusOK, views.Index(
			context.QueryParam("query"),
			query.NewQuery(params),
			cities,
		))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "5555"
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
