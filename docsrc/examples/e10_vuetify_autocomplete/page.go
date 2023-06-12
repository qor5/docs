package e10_vuetify_autocomplete

// @snippet_begin(VuetifyAutoCompleteSample)

import (
	"fmt"

	"github.com/qor5/admin/presets"
	"github.com/qor5/admin/presets/gorm2op"
	. "github.com/qor5/ui/vuetify"
	"github.com/qor5/ui/vuetifyx"
	"github.com/qor5/web"
	h "github.com/theplant/htmlgo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type (
	User struct {
		Login string
		Name  string
	}

	UserIcons struct {
		Login string `json:"text"`
		Name  string `json:"value"`
		Icon  string `json:"icon"`
	}

	Product struct {
		ID   uint `gorm:"primarykey"`
		Name string
	}
)

var (
	options = []*User{
		{Login: "sam", Name: "Sam"},
		{Login: "john", Name: "John"},
		{Login: "charles", Name: "Charles"},
	}

	iconOptions = []*UserIcons{
		{Login: "sam", Name: "Sam", Icon: "https://cdn.vuetifyjs.com/images/lists/1.jpg"},
		{Login: "john", Name: "John", Icon: "https://cdn.vuetifyjs.com/images/lists/2.jpg"},
		{Login: "charles", Name: "Charles", Icon: "https://cdn.vuetifyjs.com/images/lists/3.jpg"},
	}

	loadMoreRes   *vuetifyx.AutocompleteDataSource
	pagingRes     *vuetifyx.AutocompleteDataSource
	ExamplePreset *presets.Builder
)

func init() {
	db, err := gorm.Open(sqlite.Open("/tmp/my.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})
	db.Where("1=1").Delete(&Product{})

	for i := 1; i < 300; i++ {
		db.Create(&Product{Name: fmt.Sprintf("Product %d", i)})
	}

	ExamplePreset = presets.New()
	ExamplePreset.URIPrefix(VuetifyAutoCompletePresetPath).DataOperator(gorm2op.DataOperator(db))
	listing := ExamplePreset.Model(&Product{}).Listing()
	loadMoreRes = listing.ConfigureAutocompleteDataSource(
		&presets.AutocompleteDataSourceConfig{
			OptionValue: "ID",
			OptionText:  "Name",
			OptionIcon: func(product interface{}) string {
				return fmt.Sprintf("https://cdn.vuetifyjs.com/images/lists/%d.jpg", product.(*Product).ID%4+1)
			},
			KeywordColumns: []string{
				"Name",
			},
			PerPage: 50,
		},
		"loadMore",
	)

	pagingRes = listing.ConfigureAutocompleteDataSource(
		&presets.AutocompleteDataSourceConfig{
			OptionValue: "ID",
			OptionText:  "Name",
			OptionIcon: func(product interface{}) string {
				return fmt.Sprintf("https://cdn.vuetifyjs.com/images/lists/%d.jpg", product.(*Product).ID%4+1)
			},
			KeywordColumns: []string{
				"Name",
			},
			PerPage:  20,
			IsPaging: true,
			OrderBy:  "Name",
		},
		"paging",
	)

}

func VuetifyAutocomplete(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = VContainer(
		h.H1("Select many (default)"),
		vuetifyx.VXAutocomplete().
			Label("Load options from a list").
			Items(options).
			FieldName("Values1").
			ItemText("Name").
			ItemValue("Login"),

		h.H1("Select one"),
		vuetifyx.VXAutocomplete().
			FieldName("Values2").
			Label("Load options from a list").
			Items(options).
			ItemText("Name").
			ItemValue("Login").
			Multiple(false),

		h.H1("Has icon"),
		vuetifyx.VXAutocomplete().
			Label("Load options from a list").
			Items(iconOptions).
			HasIcon(true),

		h.H1("Load more from remote resource"),
		vuetifyx.VXAutocomplete().
			FieldName("Values2").
			Label("Load options from data source").
			SetDataSource(loadMoreRes),

		h.H1("Paging with remote resource"),
		vuetifyx.VXAutocomplete().
			FieldName("Values2").
			Label("Load options from data source").
			SetDataSource(pagingRes),

		h.H1("Sorting"),
		vuetifyx.VXAutocomplete().
			FieldName("Values2").
			Label("Load options from data source").
			Sorting(true).
			SetDataSource(pagingRes).ChipColor("red"),
	)
	return
}

var VuetifyAutocompletePB = web.Page(VuetifyAutocomplete)

const VuetifyAutoCompletePath = "/samples/vuetify-auto-complete"
const VuetifyAutoCompletePresetPath = "/samples/vuetify-auto-complete-preset"

// @snippet_end
