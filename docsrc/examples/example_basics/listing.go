package example_basics

import (
	"os"
	"time"

	"github.com/qor5/admin/presets"
	"github.com/qor5/admin/presets/gorm2op"
	"github.com/qor5/ui/vuetify"
	"github.com/qor5/web"
	h "github.com/theplant/htmlgo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	DB = setupDB()
}

func setupDB() (db *gorm.DB) {
	var err error
	db, err = gorm.Open(postgres.Open(os.Getenv("DB_PARAMS")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Logger.LogMode(logger.Info)
	err = db.AutoMigrate(
		&Post{},
		&Category{},
		&Product{},
	)
	if err != nil {
		panic(err)
	}
	return
}

const ListingSamplePath = "/samples/listing"

// @snippet_begin(PresetsListingSample)

type Post struct {
	ID        uint
	Title     string
	Body      string
	UpdatedAt time.Time
	CreatedAt time.Time
	Disabled  bool

	Status string

	CategoryID uint
}

type Category struct {
	ID   uint
	Name string

	UpdatedAt time.Time
	CreatedAt time.Time
}

func ListingSample(b *presets.Builder) {
	db := DB

	// Setup the project name, ORM and Homepage
	b.URIPrefix(ListingSamplePath).DataOperator(gorm2op.DataOperator(db))

	// Register Post into the builder
	// Use m to customize the model, Or config more models here.
	postModelBuilder := b.Model(&Post{})
	postModelBuilder.Listing("ID", "Title", "Body", "CategoryID", "VirtualField")

	postModelBuilder.Listing().Searcher = func(model interface{}, params *presets.SearchParams, ctx *web.EventContext) (r interface{}, totalCount int, err error) {
		qdb := db.Where("disabled != true")
		return gorm2op.DataOperator(qdb).Search(model, params, ctx)
	}

	rmn := postModelBuilder.Listing().RowMenu()
	rmn.RowMenuItem("Show").ComponentFunc(func(obj interface{}, id string, ctx *web.EventContext) h.HTMLComponent {
		return h.Text("Fake Show")
	})

	postModelBuilder.Listing().ActionsAsMenu(true)

	postModelBuilder.Editing().Field("CategoryID").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
		categories := []Category{}
		if err := db.Find(&categories).Error; err != nil {
			// ignore err for now
		}

		return vuetify.VAutocomplete().Chips(true).FieldName(field.Name).Label(field.Label).Value(field.Value(obj)).Items(categories).ItemText("Name").ItemValue("ID")
	})

	postModelBuilder.Listing().Field("CategoryID").Label("Category").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
		c := Category{}
		cid, _ := field.Value(obj).(uint)
		if err := db.Where("id = ?", cid).Find(&c).Error; err != nil {
			// ignore err in the example
		}
		return h.Td(h.Text(c.Name))
	})

	postModelBuilder.Listing().Field("VirtualField").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
		return h.Td(h.Text("virtual field"))
	})

	b.Model(&Category{})
	// Use m to customize the model, Or config more models here.
	return
}

// @snippet_end
