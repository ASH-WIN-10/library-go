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
    ReadStatus BOOLEAN NOT NULL
    );`

	_, err := m.DB.Exec(stmt)
	return err
}
