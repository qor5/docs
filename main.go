package docs

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	mux := Mux("/")
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9101"
	}
	// @snippet_begin(HelloWorldMainSample)
	fmt.Println("Starting docs at :" + port)
	http.Handle("/", mux)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
	// @snippet_end
}
