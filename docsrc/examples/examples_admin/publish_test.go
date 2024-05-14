package examples_admin

import (
	"net/http/httptest"
	"testing"

	"github.com/qor5/admin/v3/presets"
	"github.com/qor5/admin/v3/presets/gorm2op"
)

func TestPublish(t *testing.T) {
	pb := presets.New().DataOperator(gorm2op.DataOperator(TestDB))
	PublishExample(pb, TestDB)
	// gofixtures.Data(gofixtures.Sql(``))

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/samples/publish/products", nil)
	pb.ServeHTTP(w, r)
	t.Log(w.Body.String())
}
