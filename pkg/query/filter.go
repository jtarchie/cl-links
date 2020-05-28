package query

import (
	"github.com/jtarchie/cl-search/pkg/load"
)

func (q Query) Filter(cities load.Cities) load.Cities {
	if v, err := q.params.GetString("country"); err == nil {
		cities = cities.FilterByString(
			func(c load.City) string { return c.CountryName },
			v,
		)
	}

	if v, err := q.params.GetString("region"); err == nil {
		cities = cities.FilterByString(
			func(c load.City) string { return c.RegionName },
			v,
		)
	}

	if v, err := q.params.GetString("city"); err == nil {
		cities = cities.FilterByString(
			func(c load.City) string { return c.Name },
			v,
		)
	}

	return cities
}
