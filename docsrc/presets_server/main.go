package main

import (
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/qor5/docs/docsrc/examples/presets_mux"
	"github.com/qor5/web"
	. "github.com/theplant/htmlgo"
	"net/http"
	"os"
)

type indexMux struct {
	mux   *http.ServeMux
	paths []string
}

func (im *indexMux) Page(ctx *web.EventContext) (r web.PageResponse, err error) {
	ul := Ul()
	for _, p := range im.paths {
		ul.AppendChildren(Li(A().Href(p).Text(p)))
	}
	r.Body = ul
	return
}

func (im *indexMux) Handle(pattern string, handler http.Handler) {
	im.paths = append(im.paths, pattern)
	im.mux.Handle(pattern, handler)
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "7800"
	}

	fmt.Println("Starting docs at :" + port)
	mux := http.NewServeMux()
	var im = &indexMux{mux: http.NewServeMux()}
	presets_mux.SamplesHandler(im, "/samples")
	mux.Handle("/samples/",
		middleware.Logger(
			middleware.RequestID(
				im.mux,
			),
		),
	)
	mux.Handle("/", web.New().Page(im.Page))

	err := http.ListenAndServe(":"+port, presets_mux.Mux(mux, "/"))
	if err != nil {
		panic(err)
	}
}
