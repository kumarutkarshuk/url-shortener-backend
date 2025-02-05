package router

import (
	"urlshortener/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	// router
	r := mux.NewRouter()

	// mount routes
	r.HandleFunc("/shorten", controller.CreateShortURL).Methods("POST")
	r.HandleFunc("/{shortURL}", controller.RedirectToOriginal).Methods("GET")

	return r
}