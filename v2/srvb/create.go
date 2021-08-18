package srvb

import (
	"net/http"

	"github.com/mcaci/msdb5/v2/app/briscola"
)

func Create(w http.ResponseWriter, r *http.Request) {
	g := briscola.NewGame(briscola.WithDefaultOptions)
	if g == nil {
		http.Error(w, "game could not be created", http.StatusBadRequest)
		return
	}
	w.Write([]byte("OK"))
}

func Handler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/create", Create)
	return r
}
