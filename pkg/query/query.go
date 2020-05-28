package query

import (
	"fmt"
	"github.com/jtarchie/cl-search/pkg/load"
	"github.com/jtarchie/cl-search/pkg/parser"
	"net/url"
	"sort"
	"strconv"
)

type Query struct {
	params *parser.Params
}

func (q Query) URL(city load.City) string {
	uri, _ := url.Parse(city.URL)

	qs := url.Values{}

	if v, _ := q.params.GetString("q"); v != "" {
		qs.Add("query", v)
	}

	if v, _ := q.params.GetBoolean("bundle_duplicates"); v {
		qs.Add("bundleDuplicates", "1")
	}

	if v, _ := q.params.GetBoolean("has_image"); v {
		qs.Add("hasPic", "1")
	}

	if v, _ := q.params.GetBoolean("include_nearby"); v {
		qs.Add("searchNearby", "2")

		ids := []string{}
		for _, id := range city.NearbyCities {
			ids = append(ids, id)
		}

		sort.Strings(ids)

		for _, id := range ids {
			qs.Add("nearbyArea", id)
		}
	}

	for _, k := range q.params.Keys() {
		if v, _ := q.params.GetRange(k); v != nil {
			qs.Add(fmt.Sprintf("min_%s", k), strconv.Itoa(v.Min))
			qs.Add(fmt.Sprintf("max_%s", k), strconv.Itoa(v.Max))
		}
	}

	uri.Path = "/search/sss"
	uri.RawQuery = qs.Encode()

	return uri.String()
}

func (q Query) String() string {
	return ""
}

func NewQuery(params *parser.Params) *Query {
	return &Query{
		params: params,
	}
}
