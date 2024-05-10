package basics

import (
	"github.com/qor5/docs/v3/docsrc/examples/examples_presets"
	"github.com/qor5/docs/v3/docsrc/generated"
	"github.com/qor5/docs/v3/docsrc/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
	h "github.com/theplant/htmlgo"
)

var ManageMenu = Doc(
	Markdown(`
Menu refers to the list on the left side of the page, such as the menu of the Demo below contains Customers and Companies.
`),
	h.Br(),
	utils.Demo("", examples_presets.PresetsDetailPageCardsPath+"/customers", ""),
	Markdown(`
## Menu order
Sorting menus is very simple, use ~MenuOrder~ to sort menus as you want by **slug name** .
`),
	ch.Code(generated.MenuOrderSample).Language("go"),
	utils.DemoWithSnippetLocation("Menu Order", examples_presets.PresetsMenuOrderPath+"/books", generated.MenuOrderSampleLocation),
	Markdown(`
## Menu group and icon
~MenuGroup~ can merge multiple items into one group, as shown in the following code.

Use ~MenuIcon~ on ~ModelBuilder~ can set the item icon, and set menu group icon by ~Icon~ following ~MenuGroup~.

Icon strings can be found at <https://fonts.google.com/icons>.
`),
	ch.Code(generated.MenuGroupSample).Language("go"),
	utils.DemoWithSnippetLocation("Menu Group", examples_presets.PresetsMenuGroupPath+"/videos", generated.MenuGroupSampleLocation),
).Title("Menu").
	Slug("basics/menu")
