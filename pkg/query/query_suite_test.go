package query_test

import (
	"github.com/jtarchie/cl-search/pkg/load"
	"github.com/jtarchie/cl-search/pkg/parser"
	"github.com/jtarchie/cl-search/pkg/query"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestQuery(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Query Suite")
}

var _ = Describe("Query", func() {
	When("generating the String the represents the query", func() {
		It("generates a nice query", func() {
			params, err := parser.ParseParams(`q:"truck" include-nearby:true price:1000-2000 auto-year:1980-2000`)
			Expect(err).NotTo(HaveOccurred())

			query := query.NewQuery(params)
			qs, err := query.String()
			Expect(err).NotTo(HaveOccurred())

			Expect(qs).To(Equal(`auto_year:1980-2000 include_nearby:true price:1000-2000 q:"truck"`))
		})
	})

	When("generating the URL", func() {
		var city load.City

		BeforeEach(func() {
			city = load.City{
				URL: "https://city.craigslist.org/",
				NearbyCities: map[string]string{
					"a": "1",
					"b": "2",
				},
			}
		})

		It("uses the city as a base domain", func() {
			params, err := parser.ParseParams(`q:"truck"`)
			Expect(err).NotTo(HaveOccurred())

			query := query.NewQuery(params)
			Expect(query.URL(city)).To(Equal(`https://city.craigslist.org/search/sss?query=truck`))
		})

		It("includes nearby cities", func() {
			params, err := parser.ParseParams(`q:"truck" include-nearby:true`)
			Expect(err).NotTo(HaveOccurred())

			query := query.NewQuery(params)
			Expect(query.URL(city)).To(Equal(`https://city.craigslist.org/search/sss?nearbyArea=1&nearbyArea=2&query=truck&searchNearby=2`))
		})

		It("disable duplicates and limits images", func() {
			params, err := parser.ParseParams(`bundle-duplicates has-image`)
			Expect(err).NotTo(HaveOccurred())

			query := query.NewQuery(params)
			Expect(query.URL(city)).To(Equal(`https://city.craigslist.org/search/sss?bundleDuplicates=1&hasPic=1`))
		})

		It("handles ranges", func() {
			params, err := parser.ParseParams(`price:1000-2000 auto-year:1980-2000`)
			Expect(err).NotTo(HaveOccurred())

			query := query.NewQuery(params)
			Expect(query.URL(city)).To(Equal(`https://city.craigslist.org/search/sss?max_auto_year=2000&max_price=2000&min_auto_year=1980&min_price=1000`))
		})

		It("handles categories for searching", func() {
			params, err := parser.ParseParams(`q:"truck" category:"cta"`)
			Expect(err).NotTo(HaveOccurred())

			query := query.NewQuery(params)
			Expect(query.URL(city)).To(Equal(`https://city.craigslist.org/search/cta?query=truck`))
		})
	})
})
