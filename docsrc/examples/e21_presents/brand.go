package e21_presents

import (
	"github.com/qor5/admin/presets"
	"github.com/qor5/admin/presets/gorm2op"
	"github.com/qor5/docs/docsrc/examples/utils"
	"github.com/qor5/ui/vuetify"
	"github.com/qor5/web"
	h "github.com/theplant/htmlgo"
)

func PresetsBrand(b *presets.Builder) {
	db := utils.InitDB()
	// @snippet_begin(BrandSample)
	b.URIPrefix(PresetsBrandPath).DataOperator(gorm2op.DataOperator(db)).
		BrandFunc(func(ctx *web.EventContext) h.HTMLComponent {
			return vuetify.VCardText(
				h.H1("Admin").Style("color: red;"),
			).Class("pa-0")
		})
	// @snippet_end
	db.AutoMigrate(&utils.Page{})
	b.Model(&utils.Page{})
}

const PresetsBrandPath = "/samples/brand"
