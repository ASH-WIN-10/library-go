package main

import (
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
		app.serverError(w, r, err)
		return
	}

	data := TemplateData{
		Books: books,
	}

	app.render(w, r, http.StatusOK, data)
}

func (app *application) addBook(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	pages, err := strconv.Atoi(r.PostForm.Get("pages"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
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
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
