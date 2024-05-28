package mux_presets

import (
	"github.com/qor5/admin/v3/presets"
	"github.com/qor5/docs/v3/docsrc/examples"
	"github.com/qor5/docs/v3/docsrc/examples/examples_presets"
	"github.com/qor5/docs/v3/docsrc/examples/mux_web_vuetify"
)

func SamplesHandler(mux mux_web_vuetify.Muxer, prefix string) {
	db := examples.ExampleDB()
	addGA := mux_web_vuetify.AddGA
	// @snippet_begin(MountPresetHelloWorldSample)
	c00 := presets.New().AssetFunc(addGA)
	examples_presets.PresetsHelloWorld(c00, db)
	mux.Handle(
		examples_presets.PresetsHelloWorldPath,
		c00,
	)
	// @snippet_end

	c01 := presets.New().AssetFunc(addGA)
	examples_presets.PresetsListingCustomizationFields(c01, db)
	mux.Handle(
		examples_presets.PresetsListingCustomizationFieldsPath,
		c01,
	)

	c02 := presets.New().AssetFunc(addGA)
	examples_presets.PresetsListingCustomizationFilters(c02, db)
	mux.Handle(
		examples_presets.PresetsListingCustomizationFiltersPath,
		c02,
	)

	c03 := presets.New().AssetFunc(addGA)
	examples_presets.PresetsListingCustomizationTabs(c03, db)
	mux.Handle(
		examples_presets.PresetsListingCustomizationTabsPath,
		c03,
	)

	c04 := presets.New().AssetFunc(addGA)
	examples_presets.PresetsListingCustomizationBulkActions(c04, db)
	mux.Handle(
		examples_presets.PresetsListingCustomizationBulkActionsPath,
		c04,
	)

	c05 := presets.New().AssetFunc(addGA)
	examples_presets.PresetsEditingCustomizationDescription(c05, db)
	mux.Handle(
		examples_presets.PresetsEditingCustomizationDescriptionPath,
		c05,
	)

	c06 := presets.New().AssetFunc(addGA)
	examples_presets.PresetsEditingCustomizationFileType(c06, db)
	mux.Handle(
		examples_presets.PresetsEditingCustomizationFileTypePath,
		c06,
	)

	c07 := presets.New().AssetFunc(addGA)
	examples_presets.PresetsEditingCustomizationValidation(c07, db)
	mux.Handle(
		examples_presets.PresetsEditingCustomizationValidationPath,
		c07,
	)

	c08 := presets.New().AssetFunc(addGA)
	examples_presets.PresetsDetailPageTopNotes(c08, db)
	mux.Handle(
		examples_presets.PresetsDetailPageTopNotesPath,
		c08,
	)

	c09 := presets.New().AssetFunc(addGA)
	examples_presets.PresetsDetailPageDetails(c09, db)
	mux.Handle(
		examples_presets.PresetsDetailPageDetailsPath,
		c09,
	)

	c10 := presets.New().AssetFunc(addGA)
	examples_presets.PresetsDetailPageCards(c10, db)
	mux.Handle(
		examples_presets.PresetsDetailPageCardsPath,
		c10,
	)

	c11 := presets.New().AssetFunc(addGA)
	examples_presets.PresetsPermissions(c11, db)
	mux.Handle(
		examples_presets.PresetsPermissionsPath,
		c11,
	)

	c12 := presets.New().AssetFunc(addGA)
	examples_presets.PresetsModelBuilderExtensions(c12, db)
	mux.Handle(
		examples_presets.PresetsModelBuilderExtensionsPath,
		c12,
	)

	c13 := presets.New().AssetFunc(addGA)
	examples_presets.PresetsBasicFilter(c13, db)
	mux.Handle(
		examples_presets.PresetsBasicFilterPath,
		c13,
	)

	c14 := presets.New().AssetFunc(addGA)
	examples_presets.PresetsNotificationCenterSample(c14, db)
	mux.Handle(
		examples_presets.NotificationCenterSamplePath,
		c14,
	)

	c15 := presets.New().AssetFunc(addGA)
	examples_presets.PresetsLinkageSelectFilterItem(c15, db)
	mux.Handle(
		examples_presets.PresetsLinkageSelectFilterItemPath,
		c15,
	)

	c17 := presets.New().AssetFunc(addGA)
	examples_presets.PresetsBrandTitle(c17)
	mux.Handle(
		examples_presets.PresetsBrandTitlePath,
		c17,
	)

	c18 := presets.New().AssetFunc(addGA)
	examples_presets.PresetsBrandFunc(c18)
	mux.Handle(
		examples_presets.PresetsBrandFuncPath,
		c18,
	)

	c19 := presets.New().AssetFunc(addGA)
	examples_presets.PresetsProfile(c19)
	mux.Handle(
		examples_presets.PresetsProfilePath,
		c19,
	)

	c20 := presets.New().AssetFunc(addGA)
	examples_presets.PresetsOrderMenu(c20)
	mux.Handle(
		examples_presets.PresetsMenuOrderPath,
		c20,
	)

	c21 := presets.New().AssetFunc(addGA)
	examples_presets.PresetsGroupMenu(c21)
	mux.Handle(
		examples_presets.PresetsMenuGroupPath,
		c21,
	)

	c22 := presets.New().AssetFunc(addGA)
	examples_presets.PresetsConfirmDialog(c22, db)
	mux.Handle(
		examples_presets.PresetsConfirmDialogPath,
		c22,
	)

	c25 := presets.New().AssetFunc(addGA)
	examples_presets.PresetsEditingCustomizationTabs(c25, db)
	mux.Handle(
		examples_presets.PresetsEditingCustomizationTabsPath,
		c25,
	)

	c26 := presets.New().AssetFunc(addGA)
	examples_presets.PresetsListingCustomizationSearcher(c26, db)
	mux.Handle(
		examples_presets.PresetsListingCustomizationSearcherPath,
		c26,
	)

	c27 := presets.New().AssetFunc(addGA)
	examples_presets.PresetsDetailInlineEditDetails(c27, db)
	mux.Handle(
		examples_presets.PresetsDetailInlineEditDetailsPath,
		c27,
	)
	return
}
