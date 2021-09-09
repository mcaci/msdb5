package srvb

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mcaci/msdb5/v2/app/briscola"
	"github.com/mcaci/msdb5/v2/app/misc"
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
	i, err := g.Players().Index(notJoined)
	if err != nil {
		http.Error(w, fmt.Sprintf("cannot join game for errors in selecting not joined player: %v", err), http.StatusInternalServerError)
	}
	briscola.Register("n", g)
	json.NewEncoder(w).Encode(&struct {
		Number string `json:"number"`
	}{Number: fmt.Sprint(i + 1)})
}

func notJoined(p misc.Player) bool { return p.Name() == "" }
