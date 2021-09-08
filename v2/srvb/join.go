package srvb

import (
	"encoding/json"
	"net/http"
)

func Join(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "empty request", http.StatusBadRequest)
		return
	}
	if g == nil {
		http.Error(w, "cannot join game which is not created", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&struct {
		Number string `json:"number"`
	}{Number: "1"})
}
