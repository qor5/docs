package mux_vuetifyx

import (
	"github.com/qor5/docs/docsrc/assets"
	"github.com/qor5/docs/docsrc/examples/mux_web_vuetify"
	"github.com/qor5/docs/docsrc/examples/vuetifyx_examples"
	. "github.com/qor5/ui/vuetify"
	"github.com/qor5/ui/vuetifyx"
	"github.com/qor5/web"
	"net/http"
)

func Mux(mux *http.ServeMux, prefix string) http.Handler {

	mux.Handle("/assets/main.js",
		web.PacksHandler("text/javascript",
			JSComponentsPack(),
			vuetifyx.JSComponentsPack(),
			Vuetify(""),
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
		vuetifyx_examples.VuetifyComponentsLinkageSelectPath,
		vuetifyx_examples.VuetifyComponentsLinkageSelectPB.Wrap(mux_web_vuetify.DemoVuetifyLayout),
	)
	mux.Handle(
		vuetifyx_examples.ExpansionPanelDemoPath,
		vuetifyx_examples.ExpansionPanelDemoPB.Wrap(mux_web_vuetify.DemoVuetifyLayout),
	)
	mux.Handle(
		vuetifyx_examples.KeyInfoDemoPath,
		vuetifyx_examples.KeyInfoDemoPB.Wrap(mux_web_vuetify.DemoVuetifyLayout),
	)
	mux.Handle(
		vuetifyx_examples.FilterDemoPath,
		vuetifyx_examples.FilterDemoPB.Wrap(mux_web_vuetify.DemoVuetifyLayout),
	)
	mux.Handle(
		vuetifyx_examples.DatePickersPath,
		vuetifyx_examples.DatePickersPB.Wrap(mux_web_vuetify.DemoVuetifyLayout),
	)

	return
}
