package query

import (
	"fmt"
	"sort"

	"github.com/jtarchie/cl-search/pkg/load"
)

func (q Query) Filter(cities load.Cities) load.Cities {
	if v, err := q.params.GetString("top"); err == nil {
		switch v {
		case "nearby":
			sort.Sort(&nearbyCitySort{cities})
			cities = cities.Unique(func(c load.City) string {
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

type nearbyCitySort struct {
	c load.Cities
}

func (n *nearbyCitySort) Len() int {
	return len(n.c)
}

func (n *nearbyCitySort) Less(i, j int) bool {
	return len(n.c[i].NearbyCities) > len(n.c[j].NearbyCities)
}

func (n *nearbyCitySort) Swap(i, j int) {
	n.c[i], n.c[j] = n.c[j], n.c[i]
}

var _ sort.Interface = &nearbyCitySort{}
