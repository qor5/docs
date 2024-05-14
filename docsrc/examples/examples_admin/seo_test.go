package examples_admin

import (
	"net/http/httptest"
	"testing"

	"github.com/qor5/admin/v3/presets"
)

func TestSEOExampleBasic(t *testing.T) {
	pb := presets.New()
	SEOExampleBasic(pb, TestDB)
	// gofixtures.Data(gofixtures.Sql(``))

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", SEOExampleBasicPath+"/seo-posts", nil)
	pb.ServeHTTP(w, r)
	t.Log(w.Body.String())
}
