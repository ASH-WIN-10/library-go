package main

import (
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// file server
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// routes
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("POST /add", app.addBook)
	mux.HandleFunc("DELETE /remove/{id}", app.removeBook)
	mux.HandleFunc("PUT /update/{id}", app.updateBook)

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)
	return standard.Then(mux)
}
