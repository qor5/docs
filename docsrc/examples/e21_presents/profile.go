package e21_presents

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

			return VMenu().OffsetY(true).Children(
				h.Template().Attr("v-slot:activator", "{on, attrs}").Children(
					VList(
						VListItem(
							VListItemAvatar(
								VAvatar().Class("ml-1").Color("secondary").Size(40).Children(
									h.Span(string(name[0])).Class("white--text text-h5"),
								),
							),
							VListItemContent(
								VListItemTitle(h.Text(name)),
								h.Br(),
								VListItemSubtitle(h.Text(strings.Join(roles, ", "))),
							),
						).Class("pa-0 mb-2"),
						VListItem(
							VListItemContent(
								VListItemTitle(h.Text(account)),
							),
							VListItemIcon(
								VIcon("logout").Small(true).Attr("@click", web.Plaid().URL(logoutURL).Go()),
							),
						).Class("pa-0 my-n4 ml-1").Dense(true),
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
