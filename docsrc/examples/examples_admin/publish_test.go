package examples_admin

import (
	"net/http/httptest"
	"testing"

	"github.com/qor5/admin/v3/presets"
	"github.com/qor5/admin/v3/presets/gorm2op"
	"github.com/theplant/testenv"
	"gorm.io/gorm"
)

var TestDB *gorm.DB

func TestMain(m *testing.M) {
	env, err := testenv.New().DBEnable(true).SetUp()
	if err != nil {
		panic(err)
	}
	defer env.TearDown()
	TestDB = env.DB
	m.Run()
}

func TestPublish(t *testing.T) {
	pb := presets.New().DataOperator(gorm2op.DataOperator(TestDB))
	PublishExample(pb, TestDB)
	// gofixtures.Data(gofixtures.Sql(``))

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/samples/publish/products", nil)
	pb.ServeHTTP(w, r)
	t.Log(w.Body.String())
}
