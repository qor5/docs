package advanced_functions

import (
	"github.com/qor5/docs/docsrc/examples/examples_vuetify"
	"github.com/qor5/docs/docsrc/generated"
	"github.com/qor5/docs/docsrc/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var NavigationDrawer = Doc(
	Markdown(`
Vuetify navigation drawer provide a popup layer that show on the side of the window.

Here is one example:
`),
	ch.Code(generated.VuetifyNavigationDrawerSample).Language("go"),
	utils.Demo("Vuetify Navigation Drawer", examples_vuetify.VuetifyNavigationDrawerPath, "e15_vuetify_navigation_drawer/page.go"),
).Title("Navigation Drawer").
	Slug("vuetify-components/navigation-drawer")
