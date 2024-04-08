package mux_presets

import (
	"github.com/qor5/admin/presets"
	"github.com/qor5/docs/docsrc/examples/mux_web_vuetify"
	"github.com/qor5/docs/docsrc/examples/presets_examples"
)

func SamplesHandler(mux mux_web_vuetify.Muxer, prefix string) {

	addGA := mux_web_vuetify.AddGA
	// @snippet_begin(MountPresetHelloWorldSample)
	c00 := presets.New().AssetFunc(addGA)
	presets_examples.PresetsHelloWorld(c00)
	mux.Handle(
		presets_examples.PresetsHelloWorldPath,
		c00,
	)
	// @snippet_end

	c01 := presets.New().AssetFunc(addGA)
	presets_examples.PresetsListingCustomizationFields(c01)
	mux.Handle(
		presets_examples.PresetsListingCustomizationFieldsPath,
		c01,
	)

	c02 := presets.New().AssetFunc(addGA)
	presets_examples.PresetsListingCustomizationFilters(c02)
	mux.Handle(
		presets_examples.PresetsListingCustomizationFiltersPath,
		c02,
	)

	c03 := presets.New().AssetFunc(addGA)
	presets_examples.PresetsListingCustomizationTabs(c03)
	mux.Handle(
		presets_examples.PresetsListingCustomizationTabsPath,
		c03,
	)

	c04 := presets.New().AssetFunc(addGA)
	presets_examples.PresetsListingCustomizationBulkActions(c04)
	mux.Handle(
		presets_examples.PresetsListingCustomizationBulkActionsPath,
		c04,
	)

	c05 := presets.New().AssetFunc(addGA)
	presets_examples.PresetsEditingCustomizationDescription(c05)
	mux.Handle(
		presets_examples.PresetsEditingCustomizationDescriptionPath,
		c05,
	)

	c06 := presets.New().AssetFunc(addGA)
	presets_examples.PresetsEditingCustomizationFileType(c06)
	mux.Handle(
		presets_examples.PresetsEditingCustomizationFileTypePath,
		c06,
	)

	c07 := presets.New().AssetFunc(addGA)
	presets_examples.PresetsEditingCustomizationValidation(c07)
	mux.Handle(
		presets_examples.PresetsEditingCustomizationValidationPath,
		c07,
	)

	c08 := presets.New().AssetFunc(addGA)
	presets_examples.PresetsDetailPageTopNotes(c08)
	mux.Handle(
		presets_examples.PresetsDetailPageTopNotesPath,
		c08,
	)

	c09 := presets.New().AssetFunc(addGA)
	presets_examples.PresetsDetailPageDetails(c09)
	mux.Handle(
		presets_examples.PresetsDetailPageDetailsPath,
		c09,
	)

	c10 := presets.New().AssetFunc(addGA)
	presets_examples.PresetsDetailPageCards(c10)
	mux.Handle(
		presets_examples.PresetsDetailPageCardsPath,
		c10,
	)

	c11 := presets.New().AssetFunc(addGA)
	presets_examples.PresetsPermissions(c11)
	mux.Handle(
		presets_examples.PresetsPermissionsPath,
		c11,
	)

	c12 := presets.New().AssetFunc(addGA)
	presets_examples.PresetsModelBuilderExtensions(c12)
	mux.Handle(
		presets_examples.PresetsModelBuilderExtensionsPath,
		c12,
	)

	c13 := presets.New().AssetFunc(addGA)
	presets_examples.PresetsBasicFilter(presets_examples.DB, c13)
	mux.Handle(
		presets_examples.PresetsBasicFilterPath,
		c13,
	)

	c14 := presets.New().AssetFunc(addGA)
	presets_examples.PresetsNotificationCenterSample(c14)
	mux.Handle(
		presets_examples.NotificationCenterSamplePath,
		c14,
	)

	c15 := presets.New().AssetFunc(addGA)
	presets_examples.PresetsLinkageSelectFilterItem(c15)
	mux.Handle(
		presets_examples.PresetsLinkageSelectFilterItemPath,
		c15,
	)

	c17 := presets.New().AssetFunc(addGA)
	presets_examples.PresetsBrandTitle(c17)
	mux.Handle(
		presets_examples.PresetsBrandTitlePath,
		c17,
	)

	c18 := presets.New().AssetFunc(addGA)
	presets_examples.PresetsBrandFunc(c18)
	mux.Handle(
		presets_examples.PresetsBrandFuncPath,
		c18,
	)

	c19 := presets.New().AssetFunc(addGA)
	presets_examples.PresetsProfile(c19)
	mux.Handle(
		presets_examples.PresetsProfilePath,
		c19,
	)

	c20 := presets.New().AssetFunc(addGA)
	presets_examples.PresetsOrderMenu(c20)
	mux.Handle(
		presets_examples.PresetsMenuOrderPath,
		c20,
	)

	c21 := presets.New().AssetFunc(addGA)
	presets_examples.PresetsGroupMenu(c21)
	mux.Handle(
		presets_examples.PresetsMenuGroupPath,
		c21,
	)

	c22 := presets.New().AssetFunc(addGA)
	presets_examples.PresetsConfirmDialog(presets_examples.DB, c22)
	mux.Handle(
		presets_examples.PresetsConfirmDialogPath,
		c22,
	)

	c25 := presets.New().AssetFunc(addGA)
	presets_examples.PresetsEditingCustomizationTabs(c25)
	mux.Handle(
		presets_examples.PresetsEditingCustomizationTabsPath,
		c25,
	)

	c26 := presets.New().AssetFunc(addGA)
	presets_examples.PresetsListingCustomizationSearcher(c26)
	mux.Handle(
		presets_examples.PresetsListingCustomizationSearcherPath,
		c26,
	)
	return
}
