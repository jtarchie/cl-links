package parser_test

import (
	"github.com/jtarchie/cl-search/pkg/parser"
	. "github.com/onsi/ginkgo/extensions/table"
	"math"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestParser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Parser Suite")
}

var _ = Describe("ParseParams", func() {
	DescribeTable("parsing ranges", func(q string, min, max int) {
		params, err := parser.ParseParams(q)
		Expect(err).NotTo(HaveOccurred())

		r, err := params.GetRange("price")
		Expect(err).NotTo(HaveOccurred())
		Expect(r).NotTo(BeNil())

		Expect(r.Min).To(Equal(min))
		Expect(r.Max).To(Equal(max))
	},
		Entry("only min", "price:2000-", 2000, math.MaxInt64),
		Entry("only min with greater than", "price:>2000", 2000, math.MaxInt64),
		Entry("only max", "price:-2000", 0, 2000),
		Entry("only max with less than", "price:<2000", 0, 2000),
		Entry("with min and max range", "price:2000-3000", 2000, 3000),
	)

	DescribeTable("parsing integers", func(q string, expected int) {
		params, err := parser.ParseParams(q)
		Expect(err).NotTo(HaveOccurred())

		r, err := params.GetInteger("top")
		Expect(err).NotTo(HaveOccurred())
		Expect(r).To(Equal(expected))
	},
		Entry("find the an integer", "top:10", 10),
	)

	DescribeTable("parsing booleans", func(q string, expected bool) {
		params, err := parser.ParseParams(q)
		Expect(err).NotTo(HaveOccurred())

		r, err := params.GetBoolean("has_picture")
		Expect(err).NotTo(HaveOccurred())
		Expect(r).To(Equal(expected))
	},
		Entry("defaults to true", "has-picture", true),
		Entry("set to true", "has-picture:true", true),
		Entry("set to false", "has-picture:false", false),
	)

	DescribeTable("parsing strings", func(q string, expected string) {
		params, err := parser.ParseParams(q)
		Expect(err).NotTo(HaveOccurred())

		r, err := params.GetString("q")
		Expect(err).NotTo(HaveOccurred())
		Expect(r).To(Equal(expected))
	},
		Entry("double quotes", `q:"my name is earl"`, "my name is earl"),
		Entry("single quotes", `q:'my name is earl'`, "my name is earl"),
	)

	It("handles multiple types in one", func() {
		q := `price:2000- has-picture q:"my name is earl"`

		params, err := parser.ParseParams(q)
		Expect(err).NotTo(HaveOccurred())

		a, err := params.GetString("q")
		Expect(err).NotTo(HaveOccurred())
		Expect(a).To(Equal("my name is earl"))

		b, err := params.GetBoolean("has_picture")
		Expect(err).NotTo(HaveOccurred())
		Expect(b).To(Equal(true))

		r, err := params.GetRange("price")
		Expect(err).NotTo(HaveOccurred())
		Expect(r).NotTo(BeNil())

		Expect(r.Min).To(Equal(2000))
		Expect(r.Max).To(Equal(math.MaxInt64))
	})
})
