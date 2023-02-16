package example_basics

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
	postBuilder := b.Model(&Post{})

	// get its ListingBuilder
	listing := postBuilder.Listing()

	// Call FilterDataFunc
	listing.FilterDataFunc(func(ctx *web.EventContext) vuetifyx.FilterData {
		// Prepare filter options, it is a two dimension array: [][]string{"text", "value"}
		options := []*vuetifyx.SelectItem{
			{Text: "Draft", Value: "draft"},
			{Text: "Online", Value: "online"},
		}

		return []*vuetifyx.FilterItem{
			{
				Key:      "status",
				Label:    "Status",
				ItemType: vuetifyx.ItemTypeSelect,
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
