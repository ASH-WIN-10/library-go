package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// file server
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// routes
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("POST /add", app.addBook)

	return mux
}
