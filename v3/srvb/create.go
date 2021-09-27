package srvb

import (
	"encoding/json"
	"net/http"

	"github.com/mcaci/msdb5/v3/briscola"
)

const CreateURL = "/create"

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "empty request", http.StatusBadRequest)
		return
	}
	if g != nil {
		http.Error(w, "one game already created, cannot create more", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	var req struct {
		Name string `json:"name"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	g = briscola.NewGame(&briscola.Options{WithName: req.Name})
	json.NewEncoder(w).Encode(g)
}
