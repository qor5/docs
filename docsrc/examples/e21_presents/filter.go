package e21_presents

// @snippet_begin(FilterSample)
import (
	"github.com/qor5/admin/presets"
	"github.com/qor5/admin/presets/gorm2op"
	"github.com/qor5/ui/vuetifyx"
	"github.com/qor5/web"
)

func PresetsBasicFilter(b *presets.Builder) {
	b.URIPrefix(PresetsBasicFilterPath).
		DataOperator(gorm2op.DataOperator(DB))

	// create a ModelBuilder
	videoBuilder := b.Model(&Customer{})

	// get its ListingBuilder
	listing := videoBuilder.Listing()

	// Call FilterDataFunc
	listing.FilterDataFunc(func(ctx *web.EventContext) vuetifyx.FilterData {
		// Prepare filter options, it is a two dimension array: [][]string{"text", "value"}
		options := []*vuetifyx.SelectItem{{
			Text:  "Draft",
			Value: "draft",
		}}

		return []*vuetifyx.FilterItem{
			{
				Key:      "status",
				Label:    "Status",
				ItemType: vuetifyx.ItemTypeString,
				// %s is the condition. e.g. >, >=, =, <, <=, like，
				// ? is the value of selected option
				SQLCondition: `status %s ?`,
				Options:      options,
			},
		}
	})
}

// @snippet_end

const PresetsBasicFilterPath = "/samples/basic_filter"
