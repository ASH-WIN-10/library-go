package main

import (
	"database/sql"
	"flag"
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
	addr := flag.String("addr", ":8080", "HTTP network address")
	dsn := flag.String("dsn", "sqlite.db", "SQLite data source name")
	flag.Parse()

	db, err := openDB(*dsn)
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

	log.Printf("Starting the server at port %s", *addr)
	err = http.ListenAndServe(*addr, app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
