package models

import "database/sql"

type Tag struct {
	ID   int
	Name string
}

func (t *Tag) CreateTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS tags (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL
		)
	`)
	if err != nil {
		return err
	}
	return nil
}

func (t *Tag) Insert(db *sql.DB) error {
	_, err := db.Exec(`
		INSERT INTO tags (name)
		VALUES ($1)
	`, t.Name)
	if err != nil {
		return err
	}
	return nil
}
