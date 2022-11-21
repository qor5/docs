package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/qor5/docs/cmd/qor5/website-template/admin"
)

func main() {
	mux := admin.InitApp()

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	fmt.Println("Served at http://localhost:" + port + "/admin")

	http.Handle("/", mux)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
