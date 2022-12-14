package basics

import (
	"github.com/qor5/docs/docsrc/examples/e10_vuetify_autocomplete"
	"github.com/qor5/docs/docsrc/generated"
	"github.com/qor5/docs/docsrc/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var AutoComplete = Doc(
	Markdown(`
AutoComplete is a more advanced component that vuetify provides, We extend it
So that it can fetch remote options from an event func. here we show these examples:

- An auto complete that you can select multiple with static data
- An auto complete that you can select multiple with remote fetched dynamic data
- A static normal select component

`),
	ch.Code(generated.VuetifyAutoCompleteSample).Language("go"),
	utils.Demo("Vuetify AutoComplete", e10_vuetify_autocomplete.VuetifyAutoCompletePath, "e10_vuetify_autocomplete/page.go"),
).Title("Auto Complete").
	Slug("vuetify-components/auto-complete")
