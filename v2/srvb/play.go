package srvb

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/v2/app/briscola"
	briscolad "github.com/mcaci/msdb5/v2/dom/briscola"
)

const PlayURL = "/play"

func Play(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "empty request", http.StatusBadRequest)
		return
	}
	if g == nil {
		http.Error(w, "cannot play on game which is not created", http.StatusInternalServerError)
		return
	}
	opts, err := roundOpts(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	info := briscola.Round(opts)
	json.NewEncoder(w).Encode(info)
}

func roundOpts(r *http.Request) (*briscola.RoundOpts, error) {
	defer r.Body.Close()
	var req struct {
		Name string `json:"name"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, fmt.Errorf("could not process the request: %v", err)
	}
	return &briscola.RoundOpts{
		PlIdx:        0,
		PlHand:       nil,
		CardIdx:      0,
		PlayedCards:  nil,
		NPlayers:     2,
		BriscolaCard: briscolad.Card{Item: *card.MustID(1)}}, nil
}
