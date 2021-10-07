package srvb

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mcaci/msdb5/v3/briscola"
)

const PlayURL = "/play"

func (g *Game) Play(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "empty request", http.StatusBadRequest)
		return
	}
	if g.Game == nil {
		http.Error(w, "cannot play on game which is not created", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	var req inTest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	out, err := briscola.Play(g, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&struct {
		Pl  string `json:"player"`
		Brd string `json:"board"`
	}{Pl: out.Pl.String(), Brd: fmt.Sprintf("Cards on the board: %v", out.Brd)})
}

type inTest struct {
	G string `json:"game"`
	N string `json:"name"`
	I uint8  `json:"index"`
}

func (i inTest) Name() string { return i.N }
func (i inTest) Idx() uint8   { return i.I }
