package examples_presets

import (
	"fmt"

	"github.com/qor5/admin/v3/media"
	"github.com/qor5/admin/v3/presets"
	"github.com/qor5/admin/v3/presets/gorm2op"
	"github.com/qor5/docs/v3/docsrc/examples"
	vx "github.com/qor5/ui/v3/vuetifyx"
	"github.com/qor5/web/v3"
	h "github.com/theplant/htmlgo"
	"gorm.io/gorm"
)

func PresetsDetailInlineEditDetails(b *presets.Builder, db *gorm.DB) (
	cust *presets.ModelBuilder,
	cl *presets.ListingBuilder,
	ce *presets.EditingBuilder,
	dp *presets.DetailingBuilder,
) {
	err := db.AutoMigrate(&Customer{}, &CreditCard{}, &Note{})
	if err != nil {
		panic(err)
	}
	mediaBuilder := media.New(db)
	b.DataOperator(gorm2op.DataOperator(db)).Use(mediaBuilder)

	cust = b.Model(&Customer{})
	dp = cust.Detailing("Details").Drawer(true)
	dp.Field("Details").
		SetSwitchable(true).
		ShowComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
			cu := obj.(*Customer)
			cusID := fmt.Sprint(cu.ID)

			var termAgreed string
			if cu.TermAgreedAt != nil {
				termAgreed = cu.TermAgreedAt.Format("Jan 02,15:04 PM")
			}

			detail := vx.DetailInfo(
				vx.DetailColumn(
					vx.DetailField(vx.OptionalText(cu.Name).ZeroLabel("No Name")).Label("Name"),
					vx.DetailField(vx.OptionalText(cu.Email).ZeroLabel("No Email")).Label("Email"),
					vx.DetailField(vx.OptionalText(cusID).ZeroLabel("No ID")).Label("ID"),
					vx.DetailField(vx.OptionalText(cu.CreatedAt.Format("Jan 02,15:04 PM")).ZeroLabel("")).Label("Created"),
					vx.DetailField(vx.OptionalText(termAgreed).ZeroLabel("Not Agreed Yet")).Label("Terms Agreed"),
				).Header("ACCOUNT INFORMATION"),
				vx.DetailColumn(
					vx.DetailField(h.RawHTML(cu.Description)).Label("Description"),
				).Header("DETAILS"),
			)
			return detail
		}).
		Editing("Name", "Email", "Description", "Avatar")

	return
}

func PresetsDetailInlineEditInspectTables(b *presets.Builder, db *gorm.DB) (
	cust *presets.ModelBuilder,
	cl *presets.ListingBuilder,
	ce *presets.EditingBuilder,
	dp *presets.DetailingBuilder,
) {
	err := db.AutoMigrate(&Customer{}, &CreditCard{}, &Note{})
	if err != nil {
		panic(err)
	}
	b.DataOperator(gorm2op.DataOperator(db))

	cust = b.Model(&Customer{})
	// This should inspect Notes attributes, When it is a list, It should show a standard table in detail page
	dp = cust.Detailing("CreditCards").Drawer(true)

	return
}

func PresetsDetailInlineEditDetailsInspectShowFields(b *presets.Builder, db *gorm.DB) (
	cust *presets.ModelBuilder,
	cl *presets.ListingBuilder,
	ce *presets.EditingBuilder,
	dp *presets.DetailingBuilder,
) {
	err := db.AutoMigrate(&Customer{}, &CreditCard{}, &Note{})
	if err != nil {
		panic(err)
	}
	b.DataOperator(gorm2op.DataOperator(db))

	cust = b.Model(&Customer{})
	b.URIPrefix(examples.URLPathByFunc(PresetsDetailInlineEditDetailsInspectShowFields))
	dp = cust.Detailing("Details", "CreditCards").Drawer(true)
	dp.WrapFetchFunc(func(in presets.FetchFunc) presets.FetchFunc {
		return func(obj interface{}, id string, ctx *web.EventContext) (r interface{}, err error) {
			var cus Customer
			db.Find(&cus)

			var cc []*CreditCard
			db.Find(&cc)
			cus.CreditCards = cc
			r = cus
			return
		}
	})
	dp.Field("Details").
		Editing("Name", "Email2", "Description")

	dp.Field("Email2").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
		return h.Div().Text("abc")
	})

	ccm := b.Model(&CreditCard{}).InMenu(false)
	ccm.Editing("Number")
	l := ccm.Listing("Name")
	l.Field("Name").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
		return h.Div()
	})
	return
}
