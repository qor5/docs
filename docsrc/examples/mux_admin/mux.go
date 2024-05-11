package mux_admin

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/qor5/admin/v3/presets"
	"github.com/qor5/docs/v3/docsrc"
	"github.com/qor5/docs/v3/docsrc/assets"
	"github.com/qor5/docs/v3/docsrc/examples/examples_admin"
	"github.com/qor5/docs/v3/docsrc/examples/mux_presets"
	"github.com/qor5/docs/v3/docsrc/examples/mux_web_vuetify"
	"github.com/theplant/docgo"
)

func Mux(mux *http.ServeMux, prefix string) http.Handler {
	mux_web_vuetify.Mux(mux, prefix)
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
	mux := &mux_web_vuetify.IndexMux{Mux: http.NewServeMux()}
	mux_web_vuetify.SamplesHandler(mux, prefix)
	mux_presets.SamplesHandler(mux, prefix)
	addGA := mux_web_vuetify.AddGA

	c22 := presets.New().AssetFunc(addGA)
	examples_admin.ListingSample(c22)
	mux.Handle(
		examples_admin.ListingSamplePath,
		c22,
	)

	c23 := presets.New().AssetFunc(addGA)
	examples_admin.WorkerExampleMock(c23)
	mux.Handle(
		examples_admin.WorkerExamplePath,
		c23,
	)

	c24 := presets.New().AssetFunc(addGA)
	examples_admin.ActionWorkerExampleMock(c24)
	mux.Handle(
		examples_admin.ActionWorkerExamplePath,
		c24,
	)

	c27 := presets.New().AssetFunc(addGA)
	examples_admin.InternationalizationExample(c27)
	mux.Handle(
		examples_admin.InternationalizationExamplePath,
		c27)
	c28 := presets.New().AssetFunc(addGA)
	examples_admin.LocalizationExampleMock(c28)
	mux.Handle(
		examples_admin.LocalizationExamplePath,
		c28,
	)

	c29 := presets.New().AssetFunc(addGA)
	examples_admin.PublishExample(c29, nil)
	mux.Handle(
		examples_admin.PublishExamplePath,
		c29)

	return mux.Mux
}
