package query

import (
	"fmt"
	"github.com/jtarchie/cl-search/pkg/load"
	"sort"
)

func (q Query) Filter(cities load.Cities) load.Cities {
	if v, err := q.params.GetString("top"); err == nil {
		switch v {
		case "nearby":
			sort.Slice(cities, func(i, j int) bool {
				return len(cities[i].NearbyCities) > len(cities[j].NearbyCities)
			})
			cities = cities.Unique(func(c load.City) interface{} {
				return fmt.Sprintf("%s %s", c.CountryName, c.RegionName)
			})
		}
	}

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

	sort.Sort(cities)

	return cities
}
