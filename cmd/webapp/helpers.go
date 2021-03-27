package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}
	td.CurrentYear = time.Now().Year()
	return td
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		log.Fatal(fmt.Printf("The template %s does not exist", name))
		return
	}

	err := ts.Execute(w, app.addDefaultData(td, r))
	if err != nil {
		log.Fatal(err)
	}
}
