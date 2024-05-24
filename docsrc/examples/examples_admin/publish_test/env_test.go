package publish_test

import (
	"database/sql"
	"net/http"
	"testing"

	"github.com/qor5/admin/v3/presets"
	"github.com/qor5/admin/v3/presets/gorm2op"
	"github.com/qor5/docs/v3/docsrc/examples/examples_admin"
	"github.com/theplant/gofixtures"
	"github.com/theplant/testenv"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB
	SQLDB          *sql.DB
	PresetsBuilder *presets.Builder
)

func TestMain(m *testing.M) {
	env, err := testenv.New().DBEnable(true).SetUp()
	if err != nil {
		panic(err)
	}
	defer env.TearDown()

	DB = env.DB
	SQLDB, err = DB.DB()
	if err != nil {
		panic(err)
	}
	PresetsBuilder = presets.New().DataOperator(gorm2op.DataOperator(DB))
	examples_admin.PublishExample(PresetsBuilder, DB)

	m.Run()
}

var dataSeed = gofixtures.Data(gofixtures.Sql(`
INSERT INTO "public"."with_publish_products" ("id", "created_at", "updated_at", "deleted_at", "name", "price", "status", "online_url", "scheduled_start_at", "scheduled_end_at", "actual_start_at", "actual_end_at", "version", "version_name", "parent_version") VALUES ('6', '2024-05-22 10:21:42.871908+00', '2024-05-22 10:21:42.871908+00', NULL, 'FirstWithPublishProduct', '123', 'draft', '', NULL, NULL, NULL, NULL, '2024-05-22-v01', '2024-05-22-v01', '');
`, []string{"with_publish_products"}))

var dataEmpty = gofixtures.Data(gofixtures.Sql(``, []string{"with_publish_products"}))

type Flow struct {
	db *gorm.DB
	h  http.Handler

	// global vars
	ID string
}
