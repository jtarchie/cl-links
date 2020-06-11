package parser_test

import (
	"github.com/jtarchie/cl-search/pkg/parser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParseURL", func() {
	It("handles sub-category searches", func() {
		url := `https://denver.craigslist.org/search/msa?searchNearby=1`
		params, err := parser.ParseURL(url)
		Expect(err).NotTo(HaveOccurred())

		nearby, err := params.GetBoolean("include_nearby")
		Expect(err).NotTo(HaveOccurred())
		Expect(nearby).To(BeTrue())

		category, err := params.GetString("category")
		Expect(err).NotTo(HaveOccurred())
		Expect(category).To(Equal("msa"))

		city, err := params.GetString("city")
		Expect(err).NotTo(HaveOccurred())
		Expect(city).To(Equal("denver"))
	})

	It("handles a full query", func() {
		url := `https://denver.craigslist.org/search/msa?postedToday=1&query=strat*&searchNearby=2&nearbyArea=319&nearbyArea=210&nearbyArea=713&nearbyArea=287&nearbyArea=288&nearbyArea=315`
		params, err := parser.ParseURL(url)
		Expect(err).NotTo(HaveOccurred())

		query, err := params.GetString("q")
		Expect(err).NotTo(HaveOccurred())
		Expect(query).To(Equal("strat*"))

		nearby, err := params.GetBoolean("include_nearby")
		Expect(err).NotTo(HaveOccurred())
		Expect(nearby).To(BeTrue())

		category, err := params.GetString("category")
		Expect(err).NotTo(HaveOccurred())
		Expect(category).To(Equal("msa"))

		city, err := params.GetString("city")
		Expect(err).NotTo(HaveOccurred())
		Expect(city).To(Equal("denver"))

		postedImage, err := params.GetBoolean("posted_today")
		Expect(err).NotTo(HaveOccurred())
		Expect(postedImage).To(BeTrue())
	})
})
