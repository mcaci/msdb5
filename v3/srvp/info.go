package srvp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Info(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var data struct {
		InTurn bool `json:"inturn"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	_ = json.NewEncoder(w).Encode(&data)
	fmt.Fprint(w, data)
}
