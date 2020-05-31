package main

import (
	"fmt"
	"github.com/jtarchie/cl-search/pkg/load"
	"github.com/jtarchie/cl-search/pkg/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

//go:generate go run github.com/valyala/quicktemplate/qtc

func main() {
	cities := load.AllCities()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 9,
	}))
	e.Use(middleware.Logger())

	e.GET("/", server.Index(cities))

	port := os.Getenv("PORT")
	if port == "" {
		port = "5555"
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
