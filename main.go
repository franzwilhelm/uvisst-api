package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/franzwilhelm/uvisst-api/db"
	"github.com/franzwilhelm/uvisst-api/db/models"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/notes", GetNotes).Methods("GET")
	r.HandleFunc("/notes", AddNote).Methods("POST")
	r.Use(loggingMiddleware)
	handler := cors.Default().Handler(r)

	go func() {
		if err := http.ListenAndServe(":8081", handler); err != nil {
			log.Println(err)
		}
	}()

	if err := db.Connect("postgres", "host=localhost port=5433 user=uvisst password=uvisst123 dbname=uvisst sslmode=disable"); err != nil {
		log.Fatal(err)
	}
	if err := models.Migrate(); err != nil {
		log.Fatal(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
