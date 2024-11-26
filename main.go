package main

import (
	"flag"
	"net/http"

	"github.com/qwark97/interview/controller"
	"github.com/qwark97/interview/fetcher"
	"github.com/qwark97/interview/store"
)

func main() {
	var baseURL string
	flag.StringVar(&baseURL, "base-url", "https://api.vend.com", "Base URL of the vendor API")
	flag.Parse()

	storage := store.New()
	fetcher := fetcher.New(baseURL, http.DefaultClient)

	controller := controller.New(storage, fetcher)
	http.HandleFunc("/", controller.Handle)
	http.ListenAndServe(":8080", nil)
}
