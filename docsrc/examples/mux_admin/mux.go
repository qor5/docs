package mux_admin

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/qor5/docs/v3/docsrc"
	"github.com/qor5/docs/v3/docsrc/assets"
	"github.com/qor5/docs/v3/docsrc/examples"
	"github.com/qor5/docs/v3/docsrc/examples/examples_admin"
	"github.com/qor5/docs/v3/docsrc/examples/mux_presets"
	"github.com/qor5/docs/v3/docsrc/examples/mux_web_vuetify"
	"github.com/theplant/docgo"
)

func Mux(mux *http.ServeMux, prefix string) http.Handler {
	mux_web_vuetify.Mux(mux, prefix)

	im := &mux_web_vuetify.IndexMux{Mux: http.NewServeMux()}
	SamplesHandler(im, prefix)

	mux.Handle("/samples/",
		middleware.Logger(
			middleware.RequestID(
				im.Mux,
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

func SamplesHandler(mux examples.Muxer, prefix string) {
	mux_web_vuetify.SamplesHandler(mux, prefix)
	mux_presets.SamplesHandler(mux, prefix)

	examples.AddPresetExample(mux, examples_admin.ListingExample)
	examples.AddPresetExample(mux, examples_admin.WorkerExample)
	examples.AddPresetExample(mux, examples_admin.ActionWorkerExample)
	examples.AddPresetExample(mux, examples_admin.InternationalizationExample)
	examples.AddPresetExample(mux, examples_admin.LocalizationExample)
	examples.AddPresetExample(mux, examples_admin.PublishExample)
	examples.AddPresetExample(mux, examples_admin.SEOExampleBasic)
	examples.AddPresetExample(mux, examples_admin.ActivityExample)
}
