package examples_admin

import (
	"time"

	"github.com/qor5/admin/v3/presets"
	"github.com/qor5/admin/v3/presets/gorm2op"
	v "github.com/qor5/ui/v3/vuetify"
	"github.com/qor5/web/v3"
	h "github.com/theplant/htmlgo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var exampleDB *gorm.DB

func ExampleDB() (db *gorm.DB) {
	if exampleDB != nil {
		return exampleDB
	}
	var err error
	db, err = gorm.Open(postgres.Open(dbParamsString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Logger.LogMode(logger.Info)
	err = db.AutoMigrate(
		&Post{},
		&Category{},
		&WithPublishProduct{},
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
	db := ExampleDB()

	// Setup the project name, ORM and Homepage
	b.URIPrefix(ListingSamplePath).DataOperator(gorm2op.DataOperator(db))

	// Register Post into the builder
	// Use m to customize the model, Or config more models here.
	postModelBuilder := b.Model(&Post{})
	postModelBuilder.Listing("ID", "Title", "Body", "CategoryID", "VirtualField")

	postModelBuilder.Listing().SearchFunc(func(model interface{}, params *presets.SearchParams, ctx *web.EventContext) (r interface{}, totalCount int, err error) {
		qdb := db.Where("disabled != true")
		return gorm2op.DataOperator(qdb).Search(model, params, ctx)
	})

	rmn := postModelBuilder.Listing().RowMenu()
	rmn.RowMenuItem("Show").
		ComponentFunc(func(obj interface{}, id string, ctx *web.EventContext) h.HTMLComponent {
			return v.VListItem(
				web.Slot(
					v.VIcon("mdi-menu"),
				).Name("prepend"),
				v.VListItemTitle(
					h.Text("Show"),
				),
			)
		})
	postModelBuilder.Listing().ActionsAsMenu(true)
	postModelBuilder.Listing().Action("Action0")

	postModelBuilder.Editing().Field("CategoryID").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
		categories := []Category{}
		if err := db.Find(&categories).Error; err != nil {
			// ignore err for now
		}

		return v.VAutocomplete().
			Chips(true).
			Attr(web.VField(field.Name, field.Value(obj))...).Label(field.Label).
			Items(categories).
			ItemTitle("Name").
			ItemValue("ID")
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
