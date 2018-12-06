package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// our main function
func mymux() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", Hello).Methods("GET")
	router.HandleFunc("/start", Start).Methods("GET")
	router.HandleFunc("/echoplus/{msg}", EchoPlus).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
