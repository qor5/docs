package examples

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/qor5/admin/presets"
	"github.com/qor5/docs/docsrc"
	"github.com/qor5/docs/docsrc/assets"
	"github.com/qor5/docs/docsrc/examples/e21_presents"
	"github.com/qor5/docs/docsrc/examples/example_basics"
	"github.com/qor5/docs/docsrc/examples/web_vuetify_mux"
	"github.com/theplant/docgo"
	"net/http"
)

func Mux(mux *http.ServeMux, prefix string) http.Handler {
	web_vuetify_mux.Mux(mux, prefix)
	samplesMux := SamplesHandler(prefix)
	mux.Handle("/samples/",
		middleware.Logger(
			middleware.RequestID(
				samplesMux,
			),
		),
	)

	mux.Handle("/", docgo.New().
		MainPageTitle("QOR5 Document").
		Assets("/assets/", assets.Assets).
		DocTree(docsrc.DocTree...).
		Build(),
	)

	return mux
}

func SamplesHandler(prefix string) http.Handler {
	mux := http.NewServeMux()
	web_vuetify_mux.SamplesHandler(mux, prefix)
	addGA := web_vuetify_mux.AddGA
	// @snippet_begin(MountPresetHelloWorldSample)
	c00 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsHelloWorld(c00)
	mux.Handle(
		e21_presents.PresetsHelloWorldPath+"/",
		c00,
	)
	// @snippet_end

	c01 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsListingCustomizationFields(c01)
	mux.Handle(
		e21_presents.PresetsListingCustomizationFieldsPath+"/",
		c01,
	)

	c02 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsListingCustomizationFilters(c02)
	mux.Handle(
		e21_presents.PresetsListingCustomizationFiltersPath+"/",
		c02,
	)

	c03 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsListingCustomizationTabs(c03)
	mux.Handle(
		e21_presents.PresetsListingCustomizationTabsPath+"/",
		c03,
	)

	c04 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsListingCustomizationBulkActions(c04)
	mux.Handle(
		e21_presents.PresetsListingCustomizationBulkActionsPath+"/",
		c04,
	)

	c05 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsEditingCustomizationDescription(c05)
	mux.Handle(
		e21_presents.PresetsEditingCustomizationDescriptionPath+"/",
		c05,
	)

	c06 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsEditingCustomizationFileType(c06)
	mux.Handle(
		e21_presents.PresetsEditingCustomizationFileTypePath+"/",
		c06,
	)

	c07 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsEditingCustomizationValidation(c07)
	mux.Handle(
		e21_presents.PresetsEditingCustomizationValidationPath+"/",
		c07,
	)

	c08 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsDetailPageTopNotes(c08)
	mux.Handle(
		e21_presents.PresetsDetailPageTopNotesPath+"/",
		c08,
	)

	c09 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsDetailPageDetails(c09)
	mux.Handle(
		e21_presents.PresetsDetailPageDetailsPath+"/",
		c09,
	)

	c10 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsDetailPageCards(c10)
	mux.Handle(
		e21_presents.PresetsDetailPageCardsPath+"/",
		c10,
	)

	c11 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsPermissions(c11)
	mux.Handle(
		e21_presents.PresetsPermissionsPath+"/",
		c11,
	)

	c12 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsModelBuilderExtensions(c12)
	mux.Handle(
		e21_presents.PresetsModelBuilderExtensionsPath+"/",
		c12,
	)

	c13 := presets.New().AssetFunc(addGA)
	example_basics.PresetsBasicFilter(c13)
	mux.Handle(
		example_basics.PresetsBasicFilterPath+"/",
		c13,
	)

	c14 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsNotificationCenterSample(c14)
	mux.Handle(
		e21_presents.NotificationCenterSamplePath+"/",
		c14,
	)

	c15 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsLinkageSelectFilterItem(c15)
	mux.Handle(
		e21_presents.PresetsLinkageSelectFilterItemPath+"/",
		c15,
	)

	c16 := presets.New().AssetFunc(addGA)
	example_basics.ListingSample(c16)
	mux.Handle(
		example_basics.ListingSamplePath+"/",
		c16,
	)

	c17 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsBrandTitle(c17)
	mux.Handle(
		e21_presents.PresetsBrandTitlePath+"/",
		c17,
	)

	c18 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsBrandFunc(c18)
	mux.Handle(
		e21_presents.PresetsBrandFuncPath+"/",
		c18,
	)

	c19 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsProfile(c19)
	mux.Handle(
		e21_presents.PresetsProfilePath+"/",
		c19,
	)

	c20 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsOrderMenu(c20)
	mux.Handle(
		e21_presents.PresetsMenuOrderPath+"/",
		c20,
	)

	c21 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsGroupMenu(c21)
	mux.Handle(
		e21_presents.PresetsMenuGroupPath+"/",
		c21,
	)

	c22 := presets.New().AssetFunc(addGA)
	example_basics.PresetsConfirmDialog(c22)
	mux.Handle(
		example_basics.PresetsConfirmDialogPath+"/",
		c22,
	)

	c23 := presets.New().AssetFunc(addGA)
	example_basics.WorkerExampleMock(c23)
	mux.Handle(
		example_basics.WorkerExamplePath+"/",
		c23,
	)

	c24 := presets.New().AssetFunc(addGA)
	example_basics.ActionWorkerExampleMock(c24)
	mux.Handle(
		example_basics.ActionWorkerExamplePath+"/",
		c24,
	)

	c25 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsEditingCustomizationTabs(c25)
	mux.Handle(
		e21_presents.PresetsEditingCustomizationTabsPath+"/",
		c25,
	)

	c26 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsListingCustomizationSearcher(c26)
	mux.Handle(
		e21_presents.PresetsListingCustomizationSearcherPath+"/",
		c26,
	)

	c27 := presets.New().AssetFunc(addGA)
	example_basics.InternationalizationExample(c27)
	mux.Handle(
		example_basics.InternationalizationExamplePath+"/",
		c27)
	c28 := presets.New().AssetFunc(addGA)
	example_basics.LocalizationExampleMock(c28)
	mux.Handle(
		example_basics.LocalizationExamplePath+"/",
		c28,
	)

	c29 := presets.New().AssetFunc(addGA)
	example_basics.PublishExample(c29)
	mux.Handle(
		example_basics.PublishExamplePath+"/",
		c29)

	return mux
}
