package main

import (
	"log"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Fatal("Internal Server Error 500")
		return
	}

	app.render(w, r, "home.page.tmpl", nil)
}
