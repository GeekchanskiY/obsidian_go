package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"

	models "obsidian_go/internal/database/models"
)

// TODO: use dotenv
const (
	host     = "localhost"
	port     = 5432
	user     = "obsidian_user"
	password = "obsidian_password"
	dbname   = "obsidian_go"
)

func Connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	log.Print("Connecting to database")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Create tables
	note := &models.Note{}
	err = note.CreateTable(db)
	if err != nil {
		panic(err)
	}

	note.Title = "Hello World"
	note.Author = "This is a test note"
	note.ParentNoteID = sql.NullInt64{
		Int64: 0,
		Valid: false,
	}
	note.CreatedAt = time.Now()
	// err = note.Insert(db)
	// if err != nil {
	// 	panic(err)
	// }

	note = &models.Note{}
	err = note.Select(db, 1, note)
	if err != nil {
		panic(err)
	}
	log.Print(note.Title)

	log.Print("Connected to database")
	return nil
}