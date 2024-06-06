package handlers

import "net/http"

func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Create Note"))
}
