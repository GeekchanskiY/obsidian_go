package models

import "database/sql"

type Question struct {
	ID          int    `json:"id"`
	Question    string `json:"question"`
	Description string `json:"description"`
	TopicID     int    `json:"topic_id"`
}

func (q *Question) CreateTable(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS questions (
		id SERIAL PRIMARY KEY,
		question VARCHAR NOT NULL,
		description VARCHAR NOT NULL,
		topic_id INTEGER REFERENCES topics(id) ON DELETE CASCADE
	);
	`)
	if err != nil {
		return err
	}
	return nil
}

func (q *Question) Insert(db *sql.DB) error {
	_, err := db.Exec(`
		INSERT INTO questions (question, description, topic_id)
		VALUES ($1, $2, $3)
	`, q.Question, q.Description, q.TopicID)
	if err != nil {
		return err
	}
	return nil
}

func (q *Question) Select(db *sql.DB, id uint) error {
	err := db.QueryRow(`
		SELECT * FROM questions WHERE id = $1
	`, id).Scan(&q.ID, &q.Question, &q.Description, &q.TopicID)
	if err != nil {
		return err
	}
	return nil
}

func (q *Question) SelectAll(db *sql.DB) ([]Question, error) {

	rows, err := db.Query(`
		SELECT * FROM questions
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	questions := []Question{}
	for rows.Next() {
		var question Question
		if err := rows.Scan(&question.ID, &question.Question, &question.Description, &question.TopicID); err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}
	return questions, nil
}

func (q *Question) Delete(db *sql.DB, id uint) error {
	_, err := db.Exec(`
		DELETE FROM questions WHERE id = $1
	`, id)
	if err != nil {
		return err
	}
	return nil
}

func (q *Question) Update(db *sql.DB, id uint) error {
	_, err := db.Exec(`
		UPDATE questions SET question = $1, description = $2, topic_id = $3
		WHERE id = $4
	`, q.Question, q.Description, q.TopicID, id)
	if err != nil {
		return err
	}
	return nil
}
