package srvb

import (
	"encoding/json"
	"net/http"
)

func Join(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Empty request", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&struct {
		Number string `json:"number"`
	}{Number: "1"})
}
