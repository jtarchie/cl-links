package load

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sort"
	"strings"
)

//go:generate go run github.com/omeid/go-resources/cmd/resources -package load -declare -var=FS -output=assets.go ../../.data/processed.json

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

func AllCities() (Cities, error) {
	var cities Cities

	reader, err := FS.Open("/../../.data/processed.json")
	if err != nil {
		return nil, fmt.Errorf("cannot load file: %s", err)
	}

	contents, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("cannot load YAML reader: %s", err)
	}

	err = yaml.UnmarshalStrict(contents, &cities)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal YAML reader: %s", err)
	}

	sort.Sort(cities)

	return cities, nil
}
