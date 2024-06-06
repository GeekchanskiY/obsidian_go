package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime/trace"

	"obsidian_go/internal/database"
	rt "obsidian_go/internal/router"
)

func main() {

	// Logging file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer file.Close()

	// Logging to file and stdout
	logger_output := io.MultiWriter(file, os.Stdout)
	log.SetOutput(logger_output)

	// Trace file
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Start trace
	if err := trace.Start(f); err != nil {
		fmt.Printf("failed to start trace: %v\n", err)
		return
	}
	defer trace.Stop()

	// Connect to database
	database.Connect()

	// Start server
	r := &rt.Router{}
	r.Route(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("The Best Router!"))
	})

	r.Route(http.MethodGet, `/hello/(?P<Message>\w+)`, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello " + rt.URLParam(r, "Message")))
	})

	r.Route("GET", "/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("something bad happened!")
	})
	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	server.ListenAndServe()
}
