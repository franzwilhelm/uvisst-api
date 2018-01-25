package main

import (
	"log"
	"net/http"

	"github.com/franzwilhelm/uvisst-api/db/models"
)

// GetNotes returns all notes in db
func GetNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := models.GetAllNotes()
	if err != nil {
		response(w, 500)
		return
	}
	response(w, 200, notes)
}

// AddNote adds a note to the db
func AddNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	if err := unmarshal(r.Body, &note); err != nil {
		response(w, 400, err.Error())
		return
	}
	if note.Data == "" {
		response(w, 400, "note data required")
		return
	}
	log.Print(note)
	err := note.Create()
	if err != nil {
		response(w, 500, err.Error())
		return
	}
	response(w, 200, note)
}
