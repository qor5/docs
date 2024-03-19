package presets_mux

import (
	"fmt"
	"github.com/qor5/admin/presets"
	"github.com/qor5/docs/docsrc/assets"
	"github.com/qor5/docs/docsrc/examples/presets_examples"
	"github.com/qor5/docs/docsrc/examples/web_vuetify_mux"
	. "github.com/qor5/ui/vuetify"
	"github.com/qor5/web"
	"net/http"
	"os"
)

var coreJSTags = func() string {
	if len(os.Getenv("DEV_CORE_JS")) > 0 {
		return `
<script src='http://localhost:3100/js/chunk-vendors.js'></script>
<script src='http://localhost:3100/js/app.js'></script>
`
	}
	return `<script src='/assets/main.js'></script>`
}()

var vuetifyJSTags = func() string {
	if len(os.Getenv("DEV_VUETIFY_JS")) > 0 {
		return `
<script src='http://localhost:3080/js/chunk-vendors.js'></script>
<script src='http://localhost:3080/js/app.js'></script>
`
	}
	return `<script src='/assets/vuetify.js'></script>`
}()

// @snippet_begin(DemoVuetifyLayoutSample)
func demoVuetifyLayout(in web.PageFunc) (out web.PageFunc) {
	return func(ctx *web.EventContext) (pr web.PageResponse, err error) {

		ctx.Injector.HeadHTML(`
			<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto+Mono" async>
			<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto:300,400,500" async>
			<link href="https://cdn.jsdelivr.net/npm/@mdi/font@5.x/css/materialdesignicons.min.css" rel="stylesheet" async>
			<link rel="stylesheet" href="/assets/vuetify.css">
			<script src='/assets/vue.js'></script>
		`)

		ctx.Injector.TailHTML(fmt.Sprintf("%s %s", vuetifyJSTags, coreJSTags))
		ctx.Injector.HeadHTML(`
		<style>
			[v-cloak] {
				display: none;
			}
		</style>
		`)

		var innerPr web.PageResponse
		innerPr, err = in(ctx)
		if err != nil {
			panic(err)
		}

		pr.Body = VApp(
			VMain(
				innerPr.Body,
			),
		)

		return
	}
}

// @snippet_end

func Mux(mux *http.ServeMux, prefix string) http.Handler {

	mux.Handle("/assets/main.js",
		web.PacksHandler("text/javascript",
			web.JSComponentsPack(),
		),
	)

	mux.Handle("/assets/vue.js",
		web.PacksHandler("text/javascript",
			web.JSVueComponentsPack(),
		),
	)

	mux.Handle("/assets/vuetify.js",
		web.PacksHandler("text/javascript",
			Vuetify(""),
			JSComponentsPack(),
			// vuetifyx.JSComponentsPack(),
		),
	)

	mux.Handle("/assets/vuetify.css",
		web.PacksHandler("text/css",
			CSSComponentsPack(),
		),
	)

	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.Write(assets.Favicon)
		return
	})

	return mux
}

type muxer interface {
	Handle(pattern string, handler http.Handler)
}

func SamplesHandler(mux muxer, prefix string) {

	addGA := web_vuetify_mux.AddGA
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
