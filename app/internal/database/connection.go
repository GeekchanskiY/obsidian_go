package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/lib/pq"

	config "obsidian_go/config"
	models "obsidian_go/internal/database/models"
)

var (
	connection *sql.DB
	once       sync.Once
)

func Connect() (*sql.DB, error) {
	var err error = nil
	once.Do(func() {
		config.LoadConfig()
		var (
			host     = os.Getenv("db_host")
			port     = os.Getenv("db_port")
			user     = os.Getenv("db_user")
			password = os.Getenv("db_password")
			dbname   = os.Getenv("db_name")
		)
		log.Println("Connecting to database as user: " + os.Getenv("db_user"))
		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
		log.Print("Connecting to database")
		connection, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Printf("Failed to open database connection: %v", err)
			return
		}

		err = connection.Ping()
		if err != nil {
			log.Printf("Failed to ping database: %v", err)
			return
		}

		// Create tables
		note := &models.Note{}
		err = note.CreateTable(connection)
		if err != nil {
			log.Printf("Failed to create tables: %v", err)
			return
		}
		log.Println("Connected to database")
	},
	)
	return connection, err
}
