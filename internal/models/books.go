package models

import "database/sql"

type Book struct {
	ID         int
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
    ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
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

func (m *BookModel) All() ([]Book, error) {
	stmt := `SELECT * FROM Book`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book

		err = rows.Scan(&b.ID, &b.Title, &b.Author, &b.Pages, &b.ReadStatus)
		if err != nil {
			return nil, err
		}

		books = append(books, b)
	}

	return books, nil
}

func (m *BookModel) Delete(id int) error {
	stmt := `DELETE FROM Book WHERE ID = ?;`

	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *BookModel) Update(id int) error {
	stmt := `UPDATE Book SET ReadStatus = NOT ReadStatus WHERE ID = ?;`

	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}
