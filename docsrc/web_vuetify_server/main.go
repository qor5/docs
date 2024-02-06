package main

import (
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/qor5/docs/docsrc/examples/web_vuetify_mux"
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

	samplesMux := web_vuetify_mux.SamplesHandler(http.NewServeMux(), "/samples")
	mux.Handle("/samples/",
		middleware.Logger(
			middleware.RequestID(
				samplesMux,
			),
		),
	)

	err := http.ListenAndServe(":"+port, web_vuetify_mux.Mux(mux, "/"))
	if err != nil {
		panic(err)
	}
}
