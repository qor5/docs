package basics

import (
	"github.com/qor5/docs/v3/docsrc/examples/examples_presets"
	"github.com/qor5/docs/v3/docsrc/examples/examples_vuetifyx"
	"github.com/qor5/docs/v3/docsrc/generated"
	"github.com/qor5/docs/v3/docsrc/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var LinkageSelect = Doc(
	Markdown(`
LinkageSelect is a component for multi-level linkage select.
    `),
	ch.Code(generated.VuetifyComponentsLinkageSelect).Language("go"),
	utils.Demo("Vuetify LinkageSelect", examples_vuetifyx.VuetifyComponentsLinkageSelectPath, "e24_vuetify_components_linkage_select/page.go"),
	Markdown(`
### Filter intergation
    `),
	ch.Code(generated.LinkageSelectFilterItem).Language("go"),
	utils.Demo("LinkageSelect Filter Item", examples_presets.PresetsLinkageSelectFilterItemPath+"/addresses", "e21_presents/linkage_select_filter_item.go"),
).Title("Linkage Select").
	Slug("vuetify-components/linkage-select")
