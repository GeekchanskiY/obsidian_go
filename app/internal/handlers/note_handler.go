package handlers

import (
	"log"
	"net/http"
)

func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Println(URLParamInt(r, "id"))
	w.Write([]byte("Create Note"))
}
