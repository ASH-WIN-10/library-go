package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/ASH-WIN-10/library-go/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (app *application) addBook(w http.ResponseWriter, r *http.Request) {
	// TODO: take input from a form
	newBook := models.Book{
		Title:      "Pride and Prejudice",
		Author:     "Jane Austin",
		Pages:      279,
		ReadStatus: true,
	}

	err := app.books.Insert(newBook)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
