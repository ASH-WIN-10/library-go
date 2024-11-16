package main

import (
	"html/template"
	"log"
	"net/http"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	log.Print(err.Error(), " method=", method, " uri=", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) render(
	w http.ResponseWriter,
	r *http.Request,
	status int,
	data TemplateData,
) {
	files := []string{
		"./ui/html/index.html",
		"./ui/html/partials/card.html",
		"./ui/html/partials/form.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "index", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}
