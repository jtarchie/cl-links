package load

import (
	"sort"
	"strings"
)

//go:generate go run ../json2go.go

type City struct {
	Name         string            `yaml:"name"`
	URL          string            `yaml:"url"`
	NearbyCities map[string]string `yaml:"nearby_cities"`
	RegionName   string            `yaml:"region_name"`
	CountryName  string            `yaml:"country_name"`
}

type Cities []City

func (c Cities) FilterByString(
	column func(city City) string,
	needle string,
) Cities {
	var newCities Cities

	needle = strings.ToLower(needle)

	for _, c := range c {
		if strings.Contains(strings.ToLower(column(c)), needle) {
			newCities = append(newCities, c)
		}
	}
	return newCities
}

func (c Cities) Unique(
	duplicate func(city City) string,
) Cities {
	var newCities Cities
	duplicates := map[string]bool{}

	for _, city := range c {
		key := duplicate(city)
		if _, ok := duplicates[key]; !ok {
			duplicates[key] = true
			newCities = append(newCities, city)
		}
	}

	return newCities
}

func (c Cities) Len() int {
	return len(c)
}

func (c Cities) Less(i, j int) bool {
	if c[i].CountryName == c[j].CountryName {
		if c[i].RegionName == c[j].RegionName {
			return c[i].Name < c[j].Name
		}
		return c[i].RegionName < c[j].RegionName
	}
	return c[i].CountryName < c[j].CountryName
}

func (c Cities) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func AllCities() Cities {
	sort.Sort(allCities)

	return allCities
}
