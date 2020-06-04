package load_test

import (
	"sort"
	"testing"

	"github.com/jtarchie/cl-search/pkg/load"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLoad(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Load Suite")
}

var _ = Describe("City", func() {
	It("sorts by country, region, and name", func() {
		cities := load.Cities{
			{CountryName: "B", RegionName: "1", Name: "zzz"},
			{CountryName: "A", RegionName: "2", Name: ""},
			{CountryName: "B", RegionName: "1", Name: "aaa"},
			{CountryName: "A", RegionName: "1", Name: ""},
		}

		sort.Sort(cities)

		Expect(cities).To(Equal(
			load.Cities{
				{CountryName: "A", RegionName: "1", Name: ""},
				{CountryName: "A", RegionName: "2", Name: ""},
				{CountryName: "B", RegionName: "1", Name: "aaa"},
				{CountryName: "B", RegionName: "1", Name: "zzz"},
			},
		))
	})
})
