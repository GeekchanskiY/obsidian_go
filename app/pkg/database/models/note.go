package models

import (
	"database/sql"
	"time"
)

type Note struct {
	ID           int
	Title        string
	Author       string
	ParentNoteID sql.NullInt64
	CreatedAt    time.Time
}

func (n *Note) CreateTable(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS notes (
		id SERIAL PRIMARY KEY,
		title VARCHAR NOT NULL,
		author VARCHAR NOT NULL,
		parent_note INTEGER REFERENCES notes(id),
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	`)
	if err != nil {
		return err
	}
	return nil
}

func (n *Note) Insert(db *sql.DB) error {
	_, err := db.Exec(`
		INSERT INTO notes (title, author, parent_note, created_at)
		VALUES ($1, $2, $3, $4)
	`, n.Title, n.Author, n.ParentNoteID, n.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (n *Note) Select(db *sql.DB, id uint, note *Note) error {
	err := db.QueryRow(`
		SELECT * FROM notes WHERE id = $1
	`, id).Scan(&note.ID, &note.Title, &note.Author, &note.ParentNoteID, &note.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (n *Note) Delete(db *sql.DB, id uint) error {
	_, err := db.Exec(`
		DELETE FROM notes WHERE id = $1
	`, id)
	if err != nil {
		return err
	}
	return nil
}
