package examples_presets

import (
	"github.com/qor5/admin/v3/presets"
	"github.com/qor5/ui/v3/vuetify"
	"github.com/qor5/web/v3"
	h "github.com/theplant/htmlgo"
)

type brand struct{}

func PresetsBrandTitle(b *presets.Builder) {
	// @snippet_begin(BrandTitleSample)
	b.URIPrefix(PresetsBrandTitlePath).
		BrandTitle("QOR5 Admin")
	// @snippet_end
	b.Model(&brand{}).Listing().PageFunc(func(ctx *web.EventContext) (r web.PageResponse, err error) {
		r.Body = vuetify.VContainer()
		return
	})
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
	b.Model(&brand{}).Listing().PageFunc(func(ctx *web.EventContext) (r web.PageResponse, err error) {
		r.Body = vuetify.VContainer()
		return
	})
}

const (
	PresetsBrandTitlePath = "/samples/brand_title"
	PresetsBrandFuncPath  = "/samples/brand_func"
)
