package presets_examples

// @snippet_begin(FilterSample)
import (
	"github.com/qor5/admin/presets"
	"github.com/qor5/admin/presets/gorm2op"
	"github.com/qor5/ui/vuetifyx"
	"github.com/qor5/web"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID        uint
	Title     string
	Body      string
	UpdatedAt time.Time
	CreatedAt time.Time
	Disabled  bool

	Status string

	CategoryID uint
}

func PresetsBasicFilter(db *gorm.DB, b *presets.Builder) {
	b.URIPrefix(PresetsBasicFilterPath).
		DataOperator(gorm2op.DataOperator(db))

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
