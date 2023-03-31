package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/qor5/docs/docsrc"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8800"
	}
	// @snippet_begin(HelloWorldMainSample)
	fmt.Println("Starting docs at :" + port)
	err := http.ListenAndServe(":"+port, docsrc.Mux("/"))
	if err != nil {
		panic(err)
	}
	// @snippet_end
}
