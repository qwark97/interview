package main

import (
	"net/http"

	"github.com/qwark97/interview/controller"
	"github.com/qwark97/interview/store"
)

func main() {
	storage := store.New()

	controller := controller.New(storage)
	http.HandleFunc("/", controller.Handle)
	http.ListenAndServe(":8080", nil)
}
