package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

type application struct {
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":9990", "HTTP network address")
	flag.Parse()

	templateCache, err := newTemplateCache("./ui/tmpl/")
	if err != nil {
		log.Fatal(err)
	}

	app := application{
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	log.Printf("Starting server on %s:\n", *addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}
