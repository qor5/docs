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

// TODO: listing field display
// TODO: editing filed display
// TODO: publish events: publish/unpublish/republish
// TODO: version list dialog display and select version ...
// TODO: delete all versions
// TODO: version duplicate
// TODO: version delete to switch another...
// TODO: schedule ...
