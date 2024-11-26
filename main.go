package main

import (
	"net/http"

	"example.com/controller"
	"example.com/store"
)

func main() {
	storage := store.New()

	controller := controller.New(storage)
	http.HandleFunc("/", controller.Handle)
	http.ListenAndServe(":8080", nil)
}
