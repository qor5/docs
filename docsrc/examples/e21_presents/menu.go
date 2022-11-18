package e21_presents

import (
	"github.com/qor5/admin/presets"
	"github.com/qor5/ui/vuetify"
	"github.com/qor5/web"
	h "github.com/theplant/htmlgo"
)

type music struct{}
type video struct{}
type book struct{}

func PresetsMenu(b *presets.Builder) {
	b.URIPrefix(PresetsMenuPath)
	b.Model(&music{}).Listing().PageFunc(func(ctx *web.EventContext) (r web.PageResponse, err error) {
		r.Body = vuetify.VContainer(
			h.Div(
				h.H1("music"),
			).Class("text-center mt-8"),
		)
		return
	})
	b.Model(&video{}).Listing().PageFunc(func(ctx *web.EventContext) (r web.PageResponse, err error) {
		r.Body = vuetify.VContainer(
			h.Div(
				h.H1("video"),
			).Class("text-center mt-8"),
		)
		return
	})
	// @snippet_begin(MenuSample)
	mb := b.Model(&book{}).MenuIcon("book")

	mb.Listing().PageFunc(func(ctx *web.EventContext) (r web.PageResponse, err error) {
		r.Body = vuetify.VContainer(
			h.Div(
				h.H1("book"),
			).Class("text-center mt-8"),
		)
		return
	})

	b.MenuOrder(
		"books",
		b.MenuGroup("Media").SubItems(
			"videos",
			"musics",
		).Icon("perm_media"),
	)
	// @snippet_end
}

const PresetsMenuPath = "/samples/menu"
