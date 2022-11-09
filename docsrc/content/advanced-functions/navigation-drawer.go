package advanced_functions

import (
	"github.com/qor5/docs/docsrc/examples"
	"github.com/qor5/docs/docsrc/examples/e15_vuetify_navigation_drawer"
	"github.com/qor5/docs/docsrc/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var NavigationDrawer = Doc(
	Markdown(`
Vuetify navigation drawer provide a popup layer that show on the side of the window.

Here is one example:
`),
	ch.Code(examples.VuetifyNavigationDrawerSample).Language("go"),
	utils.Demo("Vuetify Navigation Drawer", e15_vuetify_navigation_drawer.VuetifyNavigationDrawerPath, "e15_vuetify_navigation_drawer/page.go"),
).Title("Navigation Drawer").
	Slug("vuetify-components/navigation-drawer")
