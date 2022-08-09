package main

import (
	"flag"
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	staticDir := flag.String("staticDir", "./ui/static/", "Path to static assets")

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(*staticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return mux
}
