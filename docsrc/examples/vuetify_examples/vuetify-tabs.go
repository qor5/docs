package vuetify_examples

// @snippet_begin(VuetifySnackBarsSample)
import (
	. "github.com/qor5/ui/vuetify"
	"github.com/qor5/web"
	. "github.com/theplant/htmlgo"
)

func VuetifyTabs(ctx *web.EventContext) (pr web.PageResponse, err error) {
	//	pr.Body = VContainer(
	//		RawHTML(`
	//<div>
	//      <tr>
	//        <th class="text-left">
	//          Name
	//        </th>
	//        <th class="text-left">
	//          Calories
	//        </th>
	//      </tr>
	//  </div>
	//`))
	pr.Body = VContainer(
		web.Scope(

			VCard(
				VTabs(
					VTab(
						Text("Item One"),
					).Attr("value", "one"),
					VTab(
						Text("Item Two"),
					).Attr("value", "two"),
					VTab(
						Text("Item Three"),
					).Attr("value", "three"),
				).Attr("v-model", "vars.tab").
					Attr("bg-color", "primary"),
				VCardText(
					VWindow(
						VWindowItem(
							Text("One"),
						).Attr("value", "one"),
						VWindowItem(
							Text("Two"),
						).Attr("value", "two"),
						VWindowItem(
							Text("Three"),
						).Attr("value", "three"),
					).Attr("v-model", "vars.tab"),
				),
			)).VSlot("{vars}").Init("{tab:''}"),
	)
	return
}

var VuetifyTabsPB = web.Page(VuetifyTabs)

const VuetifyTabsPath = "/samples/vuetify-tabs"

// @snippet_end
