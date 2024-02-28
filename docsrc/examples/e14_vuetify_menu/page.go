package e14_vuetify_menu

// @snippet_begin(VuetifyMenuSample)

import (
	"github.com/qor5/docs/docsrc/utils"
	. "github.com/qor5/ui/vuetify"
	"github.com/qor5/web"
	h "github.com/theplant/htmlgo"
)

type formData struct {
	EnableMessages bool
	EnableHints    bool
}

var globalFavored bool

const favoredIconPortalName = "favoredIcon"

func HelloVuetifyMenu(ctx *web.EventContext) (pr web.PageResponse, err error) {

	var fv formData
	err = ctx.UnmarshalForm(&fv)
	if err != nil {
		return
	}

	pr.Body = VContainer(
		utils.PrettyFormAsJSON(ctx),
		web.Scope(
			VMenu(
				web.Slot(
					VBtn("Menu as Popover").
						On("click", "vars.myMenuShow = true").
						Theme("dark").
						Color("indigo"),
				).Name("activator"),

				VCard(
					VList(
						VListItem(
							VListItemAction(
								web.Portal(
									favoredIcon(),
								).Name(favoredIconPortalName),
							),
						).
							PrependAvatar("https://cdn.vuetifyjs.com/images/john.jpg").
							Title("John Leider").
							Subtitle("Founder of Vuetify.js"),
					),
					VDivider(),

					VList(
						VListItem(
							VListItemAction(
								VSwitch().Color("purple").
									Attr("v-model", "locals.EnableMessages"),
							),
							VListItemTitle(h.Text("Enable messages")),
						),
						VListItem(
							VListItemAction(
								VSwitch().Color("purple").
									Attr("v-model", "locals.EnableHints"),
							),
							VListItemTitle(h.Text("Enable hints")),
						),
					),

					VCardActions(
						VSpacer(),
						VBtn("Cancel").Variant("text").
							On("click", "vars.myMenuShow = false"),
						VBtn("Save").Color("primary").
							Variant("text").OnClick("submit"),
					),
				),
			).CloseOnContentClick(false).
				Width(200).
				Offset(true).
				Attr("v-model", "vars.myMenuShow"),
		).VSlot("{ locals }").Init(h.JSONString(fv)),
	).Attr(web.InitContextVars, `{myMenuShow: false}`)

	return
}

func favoredIcon() h.HTMLComponent {
	color := ""
	if globalFavored {
		color = "red"
	}

	return VBtn("").Icon(true).Children(
		VIcon("favorite").Color(color),
	).OnClick("toggleFavored")
}

func toggleFavored(ctx *web.EventContext) (er web.EventResponse, err error) {
	globalFavored = !globalFavored
	er.UpdatePortals = append(er.UpdatePortals, &web.PortalUpdate{
		Name: favoredIconPortalName,
		Body: favoredIcon(),
	})
	return
}

func submit(ctx *web.EventContext) (er web.EventResponse, err error) {
	er.Reload = true
	er.RunScript = "vars.myMenuShow = false"
	return
}

var HelloVuetifyMenuPB = web.Page(HelloVuetifyMenu).
	EventFunc("submit", submit).
	EventFunc("toggleFavored", toggleFavored)

const HelloVuetifyMenuPath = "/samples/hello-vuetify-menu"

// @snippet_end
