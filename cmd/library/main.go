package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ASH-WIN-10/library-go/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	books *models.BookModel
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func main() {
	dsn := "sqlite.db"
	db, err := openDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := &application{
		books: models.NewBookModel(db),
	}

	// Create the Book table, Exit if it does not exist
	if err := app.books.Migrate(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting the server at port :8080")
	err = http.ListenAndServe(":8080", app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
