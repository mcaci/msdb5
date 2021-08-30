package srvb

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mcaci/msdb5/v2/app/briscola"
)

func Create(w http.ResponseWriter, r *http.Request) {
	opts, err := options(r)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	if g != nil {
		http.Error(w, "one game already created, cannot create more", http.StatusInternalServerError)
		return
	}
	g = briscola.NewGame(opts)
	json.NewEncoder(w).Encode(g)
}

func options(r *http.Request) (*briscola.Options, error) {
	if r.Body == nil {
		return briscola.WithDefaultOptions, nil
	}
	defer r.Body.Close()
	var req struct {
		Name string `json:"name"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, fmt.Errorf("could not process the request: %v", err)
	}
	return &briscola.Options{WithName: req.Name}, nil
}
