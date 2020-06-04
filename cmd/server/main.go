package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/jtarchie/cl-search/pkg/load"
	"github.com/jtarchie/cl-search/pkg/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:generate go run github.com/valyala/quicktemplate/qtc

func main() {
	cities := load.AllCities()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{}))
	e.Use(middleware.Logger())

	e.GET("/", server.Index(cities))

	port := os.Getenv("PORT")
	if port == "" {
		port = "5555"
	}

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
