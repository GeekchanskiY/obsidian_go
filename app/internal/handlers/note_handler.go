package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"obsidian_go/internal/database"
	"obsidian_go/internal/database/models"
	"time"

	"github.com/lib/pq"
)

func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {

	db, err := database.Connect()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var n models.Note
	err = json.Unmarshal(body, &n)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Validation and default values
	n.CreatedAt = time.Now()

	err = n.Insert(db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Create Note"))
}

func SelectNotesHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	n := models.Note{}
	notes, err := n.SelectAll(db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	notes_json, err := json.Marshal(notes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(notes_json))
}

func SelectNoteByIdHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	note_id, err := URLParamInt(r, "id")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Note Id"))
		return
	}
	note := &models.Note{}

	note.Select(db, uint(note_id))
	note_json, err := json.Marshal(note)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(note_json))
}

func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {

	db, err := database.Connect()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	note_id, err := URLParamInt(r, "id")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Note Id"))
		return
	}
	note := &models.Note{}
	err = note.Delete(db, uint(note_id))
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			log.Printf("psql error: %s, %s \n", err.Code, err.Message)
			if err.Code == "23503" {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Cant delete this note, because there are other notes with this parent note"))
				return
			}
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No Note Found"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete Note"))
}
