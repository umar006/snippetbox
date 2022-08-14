package main

import (
	"flag"
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	staticDir := flag.String("staticDir", "./ui/static/", "Path to static assets")

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(*staticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(mux)
}
