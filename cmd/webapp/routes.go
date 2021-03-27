package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)

	fileServer := http.FileServer(http.Dir("./ui/assets/"))
	mux.Handle("/assets/", http.StripPrefix("/assets", fileServer))

	return mux
}
