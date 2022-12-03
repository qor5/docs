package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/qor5/docs/cmd/qor5/website-template/admin"
)

func main() {
	// CMS server
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	cmsMux := admin.InitApp()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: cmsMux,
	}
	go server.ListenAndServe()
	fmt.Println("CMS Served at http://localhost:" + port + "/admin")

	// Publish server
	u, _ := url.Parse(os.Getenv("PUBLISH_URL"))
	publishPort := u.Port()
	if publishPort == "" {
		publishPort = "9001"
	}
	publishMux := http.FileServer(http.Dir(admin.PublishDir))
	server = &http.Server{
		Addr:    ":" + publishPort,
		Handler: publishMux,
	}
	go server.ListenAndServe()
	fmt.Println("Publish Served at http://localhost:" + publishPort)

	log.Fatal(server.ListenAndServe())
}
