package query

import (
	"fmt"
	"math"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/jtarchie/cl-search/pkg/load"
	"github.com/jtarchie/cl-search/pkg/parser"
)

type Query struct {
	params *parser.Params
	qs     values
}

func (q Query) URL(city load.City) string {
	uri, _ := url.Parse(city.URL)

	qs := q.qs
	qs.Reset()

	if v, _ := q.params.GetBoolean("bundle_duplicates"); v {
		qs.add("bundleDuplicates", "1")
	}

	if v, _ := q.params.GetBoolean("has_image"); v {
		qs.add("hasPic", "1")
	}

	if v, _ := q.params.GetString("q"); v != "" {
		qs.add("query", v)
	}

	if v, _ := q.params.GetBoolean("include_nearby"); v {
		qs.add("searchNearby", "2")

		ids := []string{}
		for _, id := range city.NearbyCities {
			ids = append(ids, id)
		}

		sort.Strings(ids)

		for _, id := range ids {
			qs.add("nearbyArea", id)
		}
	}

	for _, k := range q.params.Keys() {
		if v, _ := q.params.GetRange(k); v != nil {
			qs.add(fmt.Sprintf("max_%s", k), strconv.Itoa(v.Max))
			qs.add(fmt.Sprintf("min_%s", k), strconv.Itoa(v.Min))
		}
	}

	category, _ := q.params.GetString("category")
	if category == "" {
		category = "sss"
	}

	uri.Path = fmt.Sprintf("/search/%s", category)

	uri.RawQuery = qs.String()

	return uri.String()
}

func (q *Query) String() (string, error) {
	var params []string

	for _, k := range q.params.Keys() {
		if v, _ := q.params.GetString(k); v != "" {
			params = append(params, fmt.Sprintf("%s:%q", k, v))
			continue
		}

		if v, err := q.params.GetBoolean(k); err == nil {
			if v {
				params = append(params, fmt.Sprintf("%s:true", k))
			} else {
				params = append(params, fmt.Sprintf("%s:false", k))
			}

			continue
		}

		if v, err := q.params.GetInteger(k); err == nil {
			params = append(params, fmt.Sprintf("%s:%d", k, v))
			continue
		}

		if v, _ := q.params.GetRange(k); v != nil {
			if v.Max < math.MaxInt64 {
				params = append(params, fmt.Sprintf("%s:%d-%d", k, v.Min, v.Max))
				continue
			} else {
				params = append(params, fmt.Sprintf("%s:>%d", k, v.Min))
				continue
			}
		}

		return "", fmt.Errorf("could not generate query string for %s", k)
	}

	return strings.Join(params, " "), nil
}

func NewQuery(params *parser.Params) *Query {
	return &Query{
		params: params,
	}
}
