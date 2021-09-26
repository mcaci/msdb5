package srvb

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mcaci/msdb5/v2/app/briscola"
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
	opts, err := playOpts(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	info := briscola.Play(g, opts)
	json.NewEncoder(w).Encode(info)
}

type inTest struct {
	G string `json:"game"`
	N string `json:"name"`
	I uint8  `json:"index"`
}

func (i inTest) Name() string { return i.N }
func (i inTest) Idx() uint8   { return i.I }

func playOpts(r *http.Request) (*inTest, error) {
	defer r.Body.Close()
	var req inTest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, fmt.Errorf("could not process the request: %v", err)
	}
	return &req, nil
}
