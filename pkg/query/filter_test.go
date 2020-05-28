package query_test

import (
	"github.com/jtarchie/cl-search/pkg/load"
	"github.com/jtarchie/cl-search/pkg/parser"
	"github.com/jtarchie/cl-search/pkg/query"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Filter", func() {
	var (
		cities load.Cities
	)

	BeforeEach(func() {
		cities = load.Cities{
			{CountryName: "A", RegionName: "1", Name: "aaa"},
			{CountryName: "A", RegionName: "2", Name: "zzz"},
			{CountryName: "B", RegionName: "3", Name: "000"},
			{CountryName: "C", RegionName: "4", Name: "999"},
		}
	})

	parseQuery := func(qs string) *query.Query {
		params, err := parser.ParseParams(qs)
		Expect(err).NotTo(HaveOccurred())

		return query.NewQuery(params)
	}

	When("filter by country name", func() {
		It("returns only in that country", func() {
			q := parseQuery(`country:"a"`)
			Expect(q.Filter(cities)).To(Equal(load.Cities{
				{CountryName: "A", RegionName: "1", Name: "aaa"},
				{CountryName: "A", RegionName: "2", Name: "zzz"},
			}))
		})
	})

	When("filter by region name", func() {
		It("returns only in that region", func() {
			q := parseQuery(`region:"1"`)
			Expect(q.Filter(cities)).To(Equal(load.Cities{
				{CountryName: "A", RegionName: "1", Name: "aaa"},
			}))
		})
	})

	When("filter by city name", func() {
		It("returns only in that region", func() {
			q := parseQuery(`city:"000"`)
			Expect(q.Filter(cities)).To(Equal(load.Cities{
				{CountryName: "B", RegionName: "3", Name: "000"},
			}))
		})
	})
})
