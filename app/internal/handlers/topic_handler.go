package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"obsidian_go/internal/database"
	"obsidian_go/internal/database/models"
)

func CreateTopicHandler(w http.ResponseWriter, r *http.Request) {
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

	var t, test_topic models.Topic
	err = json.Unmarshal(body, &t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Validation and default values
	if t.Number < 0 {
		http.Error(w, "Cant create topic with negative number!", http.StatusBadRequest)
		return
	}

	if err := test_topic.SelectTopicByNumber(db, t.NoteID, t.Number); err == nil {
		http.Error(w, "Topic with this number already exists!", http.StatusBadRequest)
		return
	}

	if t.Number != 1 {
		if err := test_topic.SelectTopicByNumber(db, t.NoteID, t.Number-1); err == nil {
			http.Error(w, "Cant create topic with this number!", http.StatusBadRequest)
			return
		}
	}

	err = t.Insert(db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Create Topic"))
}
