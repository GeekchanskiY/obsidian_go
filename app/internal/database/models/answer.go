package models

import "database/sql"

type Answer struct {
	ID         int    `json:"id"`
	QuestionID int    `json:"question_id"`
	Answer     string `json:"answer"`
	Correct    bool   `json:"correct"`
}

func (a *Answer) CreateTable(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS answers (
		id SERIAL PRIMARY KEY,
		question_id INTEGER REFERENCES questions(id) ON DELETE CASCADE,
		answer VARCHAR NOT NULL,
		correct BOOLEAN NOT NULL
	);
	`)
	if err != nil {
		return err
	}
	return nil
}

func (a *Answer) Insert(db *sql.DB) error {
	_, err := db.Exec(`
		INSERT INTO answers (question_id, answer, correct)
		VALUES ($1, $2, $3)
	`, a.QuestionID, a.Answer, a.Correct)
	if err != nil {
		return err
	}
	return nil
}

func (a *Answer) Select(db *sql.DB, id uint) error {
	err := db.QueryRow(`
		SELECT * FROM answers WHERE id = $1
	`, id).Scan(&a.ID, &a.QuestionID, &a.Answer, &a.Correct)
	if err != nil {
		return err
	}
	return nil
}

func (a *Answer) SelectAll(db *sql.DB) ([]Answer, error) {

	rows, err := db.Query(`
		SELECT * FROM answers
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	answers := []Answer{}
	for rows.Next() {
		var answer Answer
		if err := rows.Scan(&answer.ID, &answer.QuestionID, &answer.Answer, &answer.Correct); err != nil {
			return nil, err
		}
		answers = append(answers, answer)
	}
	return answers, nil
}

func (a *Answer) Delete(db *sql.DB, id uint) error {
	_, err := db.Exec(`
		DELETE FROM answers WHERE id = $1
	`, id)
	if err != nil {
		return err
	}
	return nil
}

func (a *Answer) Update(db *sql.DB, id uint) error {
	_, err := db.Exec(`
		UPDATE answers SET question_id = $1, answer = $2, correct = $3
		WHERE id = $4
	`)
	if err != nil {
		return err
	}
	return nil
}
