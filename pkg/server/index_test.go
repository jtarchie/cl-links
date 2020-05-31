package server_test

import (
	"fmt"
	"github.com/jtarchie/cl-search/pkg/load"
	"github.com/jtarchie/cl-search/pkg/server"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"net/url"
)

var _ = Describe("Index", func() {
	var cities load.Cities

	BeforeEach(func() {
		cities = load.Cities{}
	})

	makeRequest := func(query string) (*httptest.ResponseRecorder, error) {
		e := echo.New()

		values := url.Values{}
		values.Set("query", query)

		request := httptest.NewRequest(
			http.MethodGet,
			fmt.Sprintf("/?%s", values.Encode()),
			nil,
		)
		recorder := httptest.NewRecorder()
		context := e.NewContext(request, recorder)

		index := server.Index(cities)
		return recorder, index(context)
	}

	When("passing query that is a URL", func() {
		It("is successful", func() {
			recorder, err := makeRequest("https://denver.craigslist.org/search/msa?searchNearby=1")
			Expect(err).NotTo(HaveOccurred())

			Expect(recorder.Code).To(Equal(http.StatusOK))
		})
	})

	When("passing query in ql", func() {
		It("is successful", func() {
			recorder, err := makeRequest("q:jeep")
			Expect(err).NotTo(HaveOccurred())

			Expect(recorder.Code).To(Equal(http.StatusOK))
		})
	})

	When("passing an invalid query", func() {
		It("is errors", func() {
			recorder, err := makeRequest("invalid!query")
			Expect(err).To(HaveOccurred())

			Expect(recorder.Code).To(Equal(http.StatusOK))
		})
	})
})
