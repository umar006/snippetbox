package main

import (
	"flag"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	staticDir := flag.String("staticDir", "./ui/static/", "Path to static assets")

	router := httprouter.New()

	fileServer := http.FileServer(http.Dir(*staticDir))
	router.Handler(http.MethodGet, "/static/", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodPost, "/snippet/create", app.snippetCreate)
	router.HandlerFunc(http.MethodGet, "/snippet/view/:id", app.snippetView)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
