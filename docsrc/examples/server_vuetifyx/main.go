package main

import (
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/qor5/docs/v3/docsrc/examples/mux_vuetifyx"
	"github.com/qor5/docs/v3/docsrc/examples/mux_web_vuetify"
	"github.com/qor5/web/v3"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "7800"
	}

	fmt.Println("Starting docs at :" + port)
	mux := http.NewServeMux()
	var im = &mux_web_vuetify.IndexMux{Mux: http.NewServeMux()}
	mux_vuetifyx.SamplesHandler(im, "/samples")
	mux.Handle("/samples/",
		middleware.Logger(
			middleware.RequestID(
				im.Mux,
			),
		),
	)
	mux.Handle("/", web.New().Page(im.Page))

	err := http.ListenAndServe(":"+port, mux_vuetifyx.Mux(mux, "/"))
	if err != nil {
		panic(err)
	}
}
