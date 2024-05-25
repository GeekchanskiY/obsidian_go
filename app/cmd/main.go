package main

import (
	"io"
	"log"
	"os"

	"obsidian_go/pkg/database"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	logger_output := io.MultiWriter(file, os.Stdout)
	log.SetOutput(logger_output)
	database.Connect()
	println("Hello World")
}
