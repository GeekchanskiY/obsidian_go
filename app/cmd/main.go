package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"obsidian_go/internal/database"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	logger_output := io.MultiWriter(file, os.Stdout)
	log.SetOutput(logger_output)

	database.Connect()

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "About Page")
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8181", nil)
}
