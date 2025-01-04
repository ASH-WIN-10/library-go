package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/ASH-WIN-10/library-go/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	logger *slog.Logger
	books  *models.BookModel
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
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	addr := flag.String("addr", ":8080", "HTTP network address")
	dsn := flag.String("dsn", "sqlite.db", "SQLite data source name")
	flag.Parse()

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	app := &application{
		logger: logger,
		books:  models.NewBookModel(db),
	}

	// Create the Book table, Exit if it does not exist
	if err := app.books.Migrate(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	logger.Info("starting server", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
