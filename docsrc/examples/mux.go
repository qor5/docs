package examples

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/qor5/admin/presets"
	"github.com/qor5/docs/docsrc"
	"github.com/qor5/docs/docsrc/assets"
	"github.com/qor5/docs/docsrc/examples/admin_examples"
	"github.com/qor5/docs/docsrc/examples/presents_examples"
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
	presents_examples.PresetsHelloWorld(c00)
	mux.Handle(
		presents_examples.PresetsHelloWorldPath+"/",
		c00,
	)
	// @snippet_end

	c01 := presets.New().AssetFunc(addGA)
	presents_examples.PresetsListingCustomizationFields(c01)
	mux.Handle(
		presents_examples.PresetsListingCustomizationFieldsPath+"/",
		c01,
	)

	c02 := presets.New().AssetFunc(addGA)
	presents_examples.PresetsListingCustomizationFilters(c02)
	mux.Handle(
		presents_examples.PresetsListingCustomizationFiltersPath+"/",
		c02,
	)

	c03 := presets.New().AssetFunc(addGA)
	presents_examples.PresetsListingCustomizationTabs(c03)
	mux.Handle(
		presents_examples.PresetsListingCustomizationTabsPath+"/",
		c03,
	)

	c04 := presets.New().AssetFunc(addGA)
	presents_examples.PresetsListingCustomizationBulkActions(c04)
	mux.Handle(
		presents_examples.PresetsListingCustomizationBulkActionsPath+"/",
		c04,
	)

	c05 := presets.New().AssetFunc(addGA)
	presents_examples.PresetsEditingCustomizationDescription(c05)
	mux.Handle(
		presents_examples.PresetsEditingCustomizationDescriptionPath+"/",
		c05,
	)

	c06 := presets.New().AssetFunc(addGA)
	presents_examples.PresetsEditingCustomizationFileType(c06)
	mux.Handle(
		presents_examples.PresetsEditingCustomizationFileTypePath+"/",
		c06,
	)

	c07 := presets.New().AssetFunc(addGA)
	presents_examples.PresetsEditingCustomizationValidation(c07)
	mux.Handle(
		presents_examples.PresetsEditingCustomizationValidationPath+"/",
		c07,
	)

	c08 := presets.New().AssetFunc(addGA)
	presents_examples.PresetsDetailPageTopNotes(c08)
	mux.Handle(
		presents_examples.PresetsDetailPageTopNotesPath+"/",
		c08,
	)

	c09 := presets.New().AssetFunc(addGA)
	presents_examples.PresetsDetailPageDetails(c09)
	mux.Handle(
		presents_examples.PresetsDetailPageDetailsPath+"/",
		c09,
	)

	c10 := presets.New().AssetFunc(addGA)
	presents_examples.PresetsDetailPageCards(c10)
	mux.Handle(
		presents_examples.PresetsDetailPageCardsPath+"/",
		c10,
	)

	c11 := presets.New().AssetFunc(addGA)
	presents_examples.PresetsPermissions(c11)
	mux.Handle(
		presents_examples.PresetsPermissionsPath+"/",
		c11,
	)

	c12 := presets.New().AssetFunc(addGA)
	presents_examples.PresetsModelBuilderExtensions(c12)
	mux.Handle(
		presents_examples.PresetsModelBuilderExtensionsPath+"/",
		c12,
	)

	c13 := presets.New().AssetFunc(addGA)
	admin_examples.PresetsBasicFilter(c13)
	mux.Handle(
		admin_examples.PresetsBasicFilterPath+"/",
		c13,
	)

	c14 := presets.New().AssetFunc(addGA)
	presents_examples.PresetsNotificationCenterSample(c14)
	mux.Handle(
		presents_examples.NotificationCenterSamplePath+"/",
		c14,
	)

	c15 := presets.New().AssetFunc(addGA)
	presents_examples.PresetsLinkageSelectFilterItem(c15)
	mux.Handle(
		presents_examples.PresetsLinkageSelectFilterItemPath+"/",
		c15,
	)

	c16 := presets.New().AssetFunc(addGA)
	admin_examples.ListingSample(c16)
	mux.Handle(
		admin_examples.ListingSamplePath+"/",
		c16,
	)

	c17 := presets.New().AssetFunc(addGA)
	presents_examples.PresetsBrandTitle(c17)
	mux.Handle(
		presents_examples.PresetsBrandTitlePath+"/",
		c17,
	)

	c18 := presets.New().AssetFunc(addGA)
	presents_examples.PresetsBrandFunc(c18)
	mux.Handle(
		presents_examples.PresetsBrandFuncPath+"/",
		c18,
	)

	c19 := presets.New().AssetFunc(addGA)
	presents_examples.PresetsProfile(c19)
	mux.Handle(
		presents_examples.PresetsProfilePath+"/",
		c19,
	)

	c20 := presets.New().AssetFunc(addGA)
	presents_examples.PresetsOrderMenu(c20)
	mux.Handle(
		presents_examples.PresetsMenuOrderPath+"/",
		c20,
	)

	c21 := presets.New().AssetFunc(addGA)
	presents_examples.PresetsGroupMenu(c21)
	mux.Handle(
		presents_examples.PresetsMenuGroupPath+"/",
		c21,
	)

	c22 := presets.New().AssetFunc(addGA)
	admin_examples.PresetsConfirmDialog(c22)
	mux.Handle(
		admin_examples.PresetsConfirmDialogPath+"/",
		c22,
	)

	c23 := presets.New().AssetFunc(addGA)
	admin_examples.WorkerExampleMock(c23)
	mux.Handle(
		admin_examples.WorkerExamplePath+"/",
		c23,
	)

	c24 := presets.New().AssetFunc(addGA)
	admin_examples.ActionWorkerExampleMock(c24)
	mux.Handle(
		admin_examples.ActionWorkerExamplePath+"/",
		c24,
	)

	c25 := presets.New().AssetFunc(addGA)
	presents_examples.PresetsEditingCustomizationTabs(c25)
	mux.Handle(
		presents_examples.PresetsEditingCustomizationTabsPath+"/",
		c25,
	)

	c26 := presets.New().AssetFunc(addGA)
	presents_examples.PresetsListingCustomizationSearcher(c26)
	mux.Handle(
		presents_examples.PresetsListingCustomizationSearcherPath+"/",
		c26,
	)

	c27 := presets.New().AssetFunc(addGA)
	admin_examples.InternationalizationExample(c27)
	mux.Handle(
		admin_examples.InternationalizationExamplePath+"/",
		c27)
	c28 := presets.New().AssetFunc(addGA)
	admin_examples.LocalizationExampleMock(c28)
	mux.Handle(
		admin_examples.LocalizationExamplePath+"/",
		c28,
	)

	c29 := presets.New().AssetFunc(addGA)
	admin_examples.PublishExample(c29)
	mux.Handle(
		admin_examples.PublishExamplePath+"/",
		c29)

	return mux
}
