package mux_vuetifyx

import (
	"net/http"

	"github.com/qor5/docs/v3/docsrc/assets"
	"github.com/qor5/docs/v3/docsrc/examples/examples_vuetifyx"
	"github.com/qor5/docs/v3/docsrc/examples/mux_web_vuetify"
	. "github.com/qor5/ui/v3/vuetify"
	"github.com/qor5/ui/v3/vuetifyx"
	"github.com/qor5/web/v3"
)

func Mux(mux *http.ServeMux, prefix string) http.Handler {
	mux.Handle("/assets/main.js",
		web.PacksHandler("text/javascript",
			JSComponentsPack(),
			vuetifyx.JSComponentsPack(),
			Vuetify(),
			web.JSComponentsPack(),
		),
	)

	mux.Handle("/assets/vue.js",
		web.PacksHandler("text/javascript",
			web.JSVueComponentsPack(),
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

func SamplesHandler(mux mux_web_vuetify.Muxer, prefix string) {
	mux.Handle(
		examples_vuetifyx.VuetifyComponentsLinkageSelectPath,
		examples_vuetifyx.VuetifyComponentsLinkageSelectPB.Wrap(mux_web_vuetify.DemoVuetifyLayout),
	)
	mux.Handle(
		examples_vuetifyx.ExpansionPanelDemoPath,
		examples_vuetifyx.ExpansionPanelDemoPB.Wrap(mux_web_vuetify.DemoVuetifyLayout),
	)
	mux.Handle(
		examples_vuetifyx.KeyInfoDemoPath,
		examples_vuetifyx.KeyInfoDemoPB.Wrap(mux_web_vuetify.DemoVuetifyLayout),
	)
	mux.Handle(
		examples_vuetifyx.FilterDemoPath,
		examples_vuetifyx.FilterDemoPB.Wrap(mux_web_vuetify.DemoVuetifyLayout),
	)
	mux.Handle(
		examples_vuetifyx.DatePickersPath,
		examples_vuetifyx.DatePickersPB.Wrap(mux_web_vuetify.DemoVuetifyLayout),
	)

	return
}
