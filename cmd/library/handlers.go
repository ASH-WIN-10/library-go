package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/ASH-WIN-10/library-go/internal/models"
)

type TemplateData struct {
	Book  models.Book
	Books []models.Book
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	books, err := app.books.All()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	files := []string{
		"./ui/html/index.html",
		"./ui/html/partials/card.html",
		"./ui/html/partials/form.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	data := TemplateData{
		Books: books,
	}

	err = ts.ExecuteTemplate(w, "index", data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (app *application) addBook(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	pages, err := strconv.Atoi(r.PostForm.Get("pages"))
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	readStatus := false
	if r.PostForm.Get("read") == "on" {
		readStatus = true
	}

	newBook := models.Book{
		Title:      r.PostForm.Get("title"),
		Author:     r.PostForm.Get("author"),
		Pages:      pages,
		ReadStatus: readStatus,
	}

	err = app.books.Insert(newBook)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
