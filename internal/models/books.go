package models

import "database/sql"

type Book struct {
	Title      string
	Author     string
	Pages      int
	ReadStatus bool
}

type BookModel struct {
	DB *sql.DB
}

func NewBookModel(db *sql.DB) *BookModel {
	return &BookModel{
		DB: db,
	}
}

func (m *BookModel) Migrate() error {
	stmt := `CREATE TABLE IF NOT EXISTS Book (
    Title TEXT NOT NULL,
    Author TEXT NOT NULL,
    Pages INTEGER NOT NULL,
    ReadStatus BOOLEAN NOT NULL DEFAULT FALSE
    );`

	_, err := m.DB.Exec(stmt)
	return err
}

func (m *BookModel) Insert(book Book) error {
	stmt := `INSERT INTO Book (Title, Author, Pages, ReadStatus)
    VALUES (?, ?, ?, ?);`

	_, err := m.DB.Exec(stmt, book.Title, book.Author, book.Pages, book.ReadStatus)
	if err != nil {
		return err
	}

	return nil
}
