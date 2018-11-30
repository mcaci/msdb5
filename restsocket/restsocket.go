package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Hello func
func Hello(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello restful world")
}

// EchoPlus func
func EchoPlus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var message string
	_ = json.NewDecoder(r.Body).Decode(&message)
	message = params["msg"] + "!!!"
	json.NewEncoder(w).Encode(message)
}
