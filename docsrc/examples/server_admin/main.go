package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/qor5/docs/v3/docsrc/examples/mux_admin"
	"github.com/qor5/docs/v3/docsrc/examples/mux_web_vuetify"
	"github.com/qor5/web/v3"
	"github.com/theplant/osenv"
)

var port = osenv.Get("PORT", "The port to serve on", "7800")

func main() {
	fmt.Println("Starting docs at :" + port)
	mux := http.NewServeMux()
	mux_web_vuetify.Mux(mux, "")

	im := &mux_web_vuetify.IndexMux{Mux: http.NewServeMux()}
	mux_admin.SamplesHandler(im, "/samples")
	mux.Handle("/samples/",
		middleware.Logger(
			middleware.RequestID(
				im.Mux,
			),
		),
	)
	mux.Handle("/", web.New().Page(im.Page))

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		panic(err)
	}
}
