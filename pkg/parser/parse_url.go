package parser

import (
	"fmt"
	"net/url"
	"strings"
)

func ParseURL(u string) (*Params, error) {
	fields := map[string]interface{}{}

	uri, err := url.Parse(u)
	if err != nil {
		return nil, fmt.Errorf("could not parse the URL: %s", err)
	}

	cityName := strings.Split(uri.Host, ".")[0]
	fields["city"] = cityName

	if includeNearby := uri.Query().Get("searchNearby"); includeNearby != "" {
		fields["include_nearby"] = true
	}

	paths := strings.Split(uri.Path, "/")
	if len(paths) > 1 {
		fields["category"] = paths[2]
	}

	if query := uri.Query().Get("query"); query != "" {
		fields["q"] = query
	}

	return NewParams(fields), nil
}
