package models

import (
	"database/sql"
	"log"
)

type Topic struct {
	ID     int    `json:"id"`
	NoteID int    `json:"note_id"`
	Number int    `json:"number"`
	Text   string `json:"text"`
}

func (t *Topic) CreateTable(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS topics (
		id SERIAL PRIMARY KEY,
		note_id INTEGER REFERENCES notes(id) ON DELETE CASCADE,
		number INTEGER NOT NULL,
		text VARCHAR NOT NULL
	);
	`)
	if err != nil {
		return err
	}
	return nil
}

func (t *Topic) Insert(db *sql.DB) error {
	res, err := db.Exec(`
		INSERT INTO topics (note_id, number, text)
		VALUES ($1, $2, $3)
	`, t.NoteID, t.Number, t.Text)
	log.Println(res)
	if err != nil {
		return err
	}
	return nil
}

func (t *Topic) Select(db *sql.DB, id uint) error {
	err := db.QueryRow(`
		SELECT * FROM topics WHERE id = $1
	`, id).Scan(&t.ID, &t.NoteID, &t.Number, &t.Text)
	if err != nil {
		return err
	}
	return nil
}

func (t *Topic) SelectTopicByNumber(db *sql.DB, note_id int, number int) error {
	err := db.QueryRow(`
		SELECT * FROM topics WHERE note_id = $1 AND number = $2
	`, note_id, number).Scan(&t.ID, &t.NoteID, &t.Number, &t.Text)
	if err != nil {
		return err
	}
	return nil
}

func (t *Topic) SelectAll(db *sql.DB) ([]Topic, error) {
	rows, err := db.Query(`
		SELECT * FROM topics
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var topics []Topic
	for rows.Next() {
		var topic Topic
		if err := rows.Scan(&topic.ID, &topic.NoteID, &topic.Number, &topic.Text); err != nil {
			return nil, err
		}
		topics = append(topics, topic)
	}
	return topics, nil
}

func (t *Topic) Delete(db *sql.DB, id uint) error {
	_, err := db.Exec(`
		DELETE FROM topics WHERE id = $1
	`, id)
	if err != nil {
		return err
	}
	return nil
}

func (t *Topic) Update(db *sql.DB, id uint) error {
	_, err := db.Exec(`
		UPDATE topics SET note_id = $1, number = $2, text = $3
		WHERE id = $4
	`, t.NoteID, t.Number, t.Text, id)
	if err != nil {
		return err
	}
	return nil
}
