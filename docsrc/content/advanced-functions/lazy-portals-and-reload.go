package advanced_functions

import (
	"github.com/qor5/docs/docsrc/examples/e17_hello_lazy_portals_and_reload"
	"github.com/qor5/docs/docsrc/generated"
	"github.com/qor5/docs/docsrc/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var LazyPortalsAndReload = Doc(
	Markdown(`
Use ~web.Portal().Loader(web.POST().EventFunc("menuItems")).Name("menuContent")~ to put a portal place holder inside a part of html, and it will load specified event func's response body inside the place holder after the main page is rendered in a separate AJAX request. Later in an event func, you could also use ~r.ReloadPortals = []string{"menuContent"}~ to reload the portal.
`),
	ch.Code(generated.LazyPortalsAndReloadSample).Language("go"),
	utils.Demo("Lazy Portals", e17_hello_lazy_portals_and_reload.LazyPortalsAndReloadPath, "e17_hello_lazy_portals_and_reload/page.go"),
).Title("Lazy Portals").
	Slug("vuetify-components/lazy-portals")
