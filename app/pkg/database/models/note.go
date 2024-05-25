package models

import "database/sql"

type Note struct {
	ID      int
	Title   string
	Content string
}

func (n *Note) CreateTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS notes (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			content TEXT NOT NULL
		)
	`)
	if err != nil {
		return err
	}
	return nil
}

func (n *Note) Insert(db *sql.DB) error {
	_, err := db.Exec(`
		INSERT INTO notes (title, content)
		VALUES ($1, $2)
	`, n.Title, n.Content)
	if err != nil {
		return err
	}
	return nil
}
