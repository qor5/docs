package presets_examples

import (
	"strings"

	"github.com/qor5/admin/presets"
	. "github.com/qor5/ui/vuetify"
	"github.com/qor5/web"
	h "github.com/theplant/htmlgo"
)

func PresetsProfile(b *presets.Builder) {
	// @snippet_begin(ProfileSample)
	b.URIPrefix(PresetsProfilePath).BrandTitle("Admin").
		ProfileFunc(func(ctx *web.EventContext) h.HTMLComponent {
			// Demo
			logoutURL := "."
			name := "QOR5"
			account := "hello@getqor.com"
			roles := []string{"Developer"}

			return VMenu().Children(
				h.Template().Attr("v-slot:activator", "{isActive,props}").Children(
					VList(
						VListItem(
							h.Template().Attr("v-slot:prepend").Children(
								VAvatar().Class("ml-1").Color("secondary").Size(40).Children(
									h.Span(string(name[0])).Class("white--text text-h5")),
							),
						).Title(name).Subtitle(strings.Join(roles, ", ")).Class("pa-0 mb-2"),
						VListItem(
							VIcon("logout").Size("small").Attr("@click", web.Plaid().URL(logoutURL).Go()),
						).Title(account).Class("pa-0 my-n4 ml-1"),
					).Class("pa-0 ma-n4"),
				),
			)
		})
	// @snippet_end
	b.Model(&brand{}).Listing().PageFunc(func(ctx *web.EventContext) (r web.PageResponse, err error) {
		r.Body = VContainer()
		return
	})
}

const PresetsProfilePath = "/samples/profile"
