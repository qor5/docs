package examples_admin

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/qor5/admin/v3/presets"
	"github.com/qor5/web/v3/multipartestutils"
)

func TestSEOExampleBasic(t *testing.T) {
	pb := presets.New()
	SEOExampleBasic(pb, TestDB)
	// gofixtures.Data(gofixtures.Sql(``))

	cases := []multipartestutils.TestCase{
		{
			Name:  "Index Page",
			Debug: true,
			ReqFunc: func() *http.Request {
				return httptest.NewRequest("GET", "/seo-posts", nil)
			},
			ExpectPageBodyContainsInOrder: []string{"Seoposts"},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			multipartestutils.RunCase(t, c, pb)
		})
	}
}
