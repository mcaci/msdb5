package srvb

import (
	"encoding/json"
	"net/http"

	"github.com/mcaci/msdb5/v2/app/briscola"
)

func Create(w http.ResponseWriter, r *http.Request) {
	opts, err := options(r)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	g := briscola.NewGame(opts)
	if g == nil {
		http.Error(w, "game could not be created", http.StatusBadRequest)
		return
	}
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
		return nil, err
	}
	return &briscola.Options{WithName: req.Name}, nil

}

func Handler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/create", Create)
	return r
}
