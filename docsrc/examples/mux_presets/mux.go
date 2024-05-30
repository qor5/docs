package mux_presets

import (
	"fmt"

	"github.com/qor5/admin/v3/presets"
	"github.com/qor5/docs/v3/docsrc/examples"
	"github.com/qor5/docs/v3/docsrc/examples/examples_presets"
	"gorm.io/gorm"
)

func SamplesHandler(mux examples.Muxer, prefix string) {
	db := examples.ExampleDB()
	addExample(mux, db, examples_presets.PresetsHelloWorld)
	addExample(mux, db, examples_presets.PresetsListingCustomizationFields)
	addExample(mux, db, examples_presets.PresetsListingCustomizationFilters)
	addExample(mux, db, examples_presets.PresetsListingCustomizationTabs)
	addExample(mux, db, examples_presets.PresetsListingCustomizationBulkActions)
	addExample(mux, db, examples_presets.PresetsEditingCustomizationDescription)
	addExample(mux, db, examples_presets.PresetsEditingCustomizationFileType)
	addExample(mux, db, examples_presets.PresetsEditingCustomizationValidation)
	addExample(mux, db, examples_presets.PresetsDetailPageTopNotes)
	addExample(mux, db, examples_presets.PresetsDetailPageDetails)
	addExample(mux, db, examples_presets.PresetsDetailPageCards)
	addExample(mux, db, examples_presets.PresetsPermissions)
	addExample(mux, db, examples_presets.PresetsModelBuilderExtensions)
	addExample(mux, db, examples_presets.PresetsBasicFilter)
	addExample(mux, db, examples_presets.PresetsNotificationCenterSample)
	addExample(mux, db, examples_presets.PresetsLinkageSelectFilterItem)
	addExample(mux, db, examples_presets.PresetsBrandTitle)
	addExample(mux, db, examples_presets.PresetsBrandFunc)
	addExample(mux, db, examples_presets.PresetsProfile)
	addExample(mux, db, examples_presets.PresetsOrderMenu)
	addExample(mux, db, examples_presets.PresetsGroupMenu)
	addExample(mux, db, examples_presets.PresetsConfirmDialog)
	addExample(mux, db, examples_presets.PresetsEditingCustomizationTabs)
	addExample(mux, db, examples_presets.PresetsListingCustomizationSearcher)
	addExample(mux, db, examples_presets.PresetsDetailInlineEditDetails)
	addExample(mux, db, examples_presets.PresetsDetailInlineEditInspectTables)
	return
}

type exampleFunc func(b *presets.Builder, db *gorm.DB) (
	cust *presets.ModelBuilder,
	cl *presets.ListingBuilder,
	ce *presets.EditingBuilder,
	dp *presets.DetailingBuilder,
)

func addExample(mux examples.Muxer, db *gorm.DB, f exampleFunc) {
	path := examples.URLPathByFunc(f)
	p := presets.New().AssetFunc(examples.AddGA).URIPrefix(path)
	f(p, db)
	fmt.Println("Example mounting at: ", path)
	mux.Handle(
		path,
		p,
	)
}
