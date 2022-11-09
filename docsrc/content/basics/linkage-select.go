package basics

import (
	"github.com/qor5/docs/docsrc/examples"
	"github.com/qor5/docs/docsrc/examples/e21_presents"
	"github.com/qor5/docs/docsrc/examples/e24_vuetify_components_linkage_select"
	"github.com/qor5/docs/docsrc/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var LinkageSelect = Doc(
	Markdown(`
LinkageSelect is a component for multi-level linkage select.
    `),
	ch.Code(examples.VuetifyComponentsLinkageSelect).Language("go"),
	utils.Demo("Vuetify LinkageSelect", e24_vuetify_components_linkage_select.VuetifyComponentsLinkageSelectPath, "e24_vuetify_components_linkage_select/page.go"),
	Markdown(`
### Filter intergation
    `),
	ch.Code(examples.LinkageSelectFilterItem).Language("go"),
	utils.Demo("LinkageSelect Filter Item", e21_presents.PresetsLinkageSelectFilterItemPath+"/addresses", "e21_presents/linkage_select_filter_item.go"),
).Title("Linkage Select").
	Slug("vuetify-components/linkage-select")
