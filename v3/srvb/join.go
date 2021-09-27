package srvb

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mcaci/msdb5/v3/briscola"
)

const JoinURL = "/join"

func Join(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "empty request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	var req struct {
		Name string `json:"name"`
		Game string `json:"game"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	switch {
	case err != nil:
		http.Error(w, fmt.Sprintf("could not process the request: %v", err), http.StatusBadRequest)
		return
	case req.Name == "":
		http.Error(w, "no player name was given", http.StatusInternalServerError)
		return
	case req.Game == "":
		http.Error(w, "no game name was given", http.StatusInternalServerError)
		return
	case g == nil:
		http.Error(w, "cannot join game which is not created", http.StatusInternalServerError)
		return
	case g.Name != req.Game:
		http.Error(w, "cannot join game with different name", http.StatusInternalServerError)
		return
	}
	err = briscola.Register(req.Name, g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	i, err := g.Players().Index(func(p briscola.Player) bool { return p.Name() == req.Name })
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&struct {
		Number string `json:"number"`
	}{Number: fmt.Sprint(1 + i)})
}
