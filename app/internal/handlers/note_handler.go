package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"obsidian_go/internal/database"
	"obsidian_go/internal/database/models"
)

func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Println(URLParamInt(r, "id"))
	w.Write([]byte("Create Note"))
}

func SelectNotesHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Select Notes"))
}

func SelectNoteByIdHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error 1 "))
	}

	note_id, err := URLParamInt(r, "id")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Note Id"))
	}
	note := &models.Note{}

	note.Select(db, uint(note_id), note)
	note_json, err := json.Marshal(note)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error 2"))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(note_json))
}
