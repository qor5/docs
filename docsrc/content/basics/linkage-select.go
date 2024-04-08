package basics

import (
	"github.com/qor5/docs/docsrc/examples/presets_examples"
	"github.com/qor5/docs/docsrc/examples/vuetifyx_examples"
	"github.com/qor5/docs/docsrc/generated"
	"github.com/qor5/docs/docsrc/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var LinkageSelect = Doc(
	Markdown(`
LinkageSelect is a component for multi-level linkage select.
    `),
	ch.Code(generated.VuetifyComponentsLinkageSelect).Language("go"),
	utils.Demo("Vuetify LinkageSelect", vuetifyx_examples.VuetifyComponentsLinkageSelectPath, "e24_vuetify_components_linkage_select/page.go"),
	Markdown(`
### Filter intergation
    `),
	ch.Code(generated.LinkageSelectFilterItem).Language("go"),
	utils.Demo("LinkageSelect Filter Item", presets_examples.PresetsLinkageSelectFilterItemPath+"/addresses", "e21_presents/linkage_select_filter_item.go"),
).Title("Linkage Select").
	Slug("vuetify-components/linkage-select")
