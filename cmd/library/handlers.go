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
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	author := r.PostForm.Get("author")
	readStatus := r.PostForm.Get("read") == "on"

	pages, err := strconv.Atoi(r.PostForm.Get("pages"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	if title == "" || author == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	newBook := models.Book{
		Title:      title,
		Author:     author,
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
