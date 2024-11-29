package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/generate",generateHandler).Methods("POST")
	r.HandleFunc("/api/retrieve/{id}",retrieveHandler).Methods("GET")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}