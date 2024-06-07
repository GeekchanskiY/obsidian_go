package models

import (
	"database/sql"
	"time"
)

type Note struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Author       string    `json:"author"`
	ParentNoteID NullInt64 `json:"parent_note_id"`
	CreatedAt    time.Time `json:"created_at"`
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

func (note *Note) Select(db *sql.DB, id uint) error {
	err := db.QueryRow(`
		SELECT * FROM notes WHERE id = $1
	`, id).Scan(&note.ID, &note.Title, &note.Author, &note.ParentNoteID, &note.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (note *Note) SelectAll(db *sql.DB) ([]Note, error) {

	rows, err := db.Query(`
		SELECT * FROM notes
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var notes []Note
	for rows.Next() {
		var note Note
		if err := rows.Scan(&note.ID, &note.Title, &note.Author, &note.ParentNoteID, &note.CreatedAt); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
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

func (n *Note) Update(db *sql.DB, id uint) error {
	_, err := db.Exec(`
		UPDATE notes SET title = $1, author = $2, parent_note = $3
		WHERE id = $4
	`, n.Title, n.Author, n.ParentNoteID, id)
	if err != nil {
		return err
	}
	return nil
}
