package e21_presents

import (
	"github.com/qor5/admin/presets"
	"github.com/qor5/ui/vuetify"
	"github.com/qor5/web"
	h "github.com/theplant/htmlgo"
)

func PresetsBrandTitle(b *presets.Builder) {
	// @snippet_begin(BrandTitleSample)
	b.URIPrefix(PresetsBrandTitlePath).
		BrandTitle("QOR5 Admin")
	// @snippet_end
}

func PresetsBrandFunc(b *presets.Builder) {
	// @snippet_begin(BrandFuncSample)
	b.URIPrefix(PresetsBrandFuncPath).
		BrandFunc(func(ctx *web.EventContext) h.HTMLComponent {
			return vuetify.VCardText(
				h.H1("Admin").Style("color: red;"),
			).Class("pa-0")
		})
	// @snippet_end
}

const PresetsBrandTitlePath = "/samples/brand_title"
const PresetsBrandFuncPath = "/samples/brand_func"
