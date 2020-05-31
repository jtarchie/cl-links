package server

import (
	"fmt"
	"github.com/jtarchie/cl-search/cmd/server/views"
	"github.com/jtarchie/cl-search/pkg/load"
	"github.com/jtarchie/cl-search/pkg/parser"
	"github.com/jtarchie/cl-search/pkg/query"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Index(cities load.Cities) func(context echo.Context) error {
	return func(context echo.Context) error {
		params, err := parser.ParseParams(context.QueryParam("query"))
		if err != nil {
			return fmt.Errorf("could not parse query: %s", err)
		}

		q := query.NewQuery(params)

		qs, err := q.String()
		if err != nil {
			return fmt.Errorf("cannot generate query string: %s", err)
		}

		return context.HTML(http.StatusOK, views.Index(
			qs,
			q,
			q.Filter(cities),
		))
	}
}
