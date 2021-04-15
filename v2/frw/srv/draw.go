package srv

import (
	"bytes"
	"net/http"

	"github.com/mcaci/ita-cards/set"
)

var (
	cards       = set.Deck()
	currentBody = [][]byte{}
)

func Draw(w http.ResponseWriter, r *http.Request) {
	currentBody = append(currentBody, []byte(cards.Top().String()))
	err := start.Execute(w, &struct {
		Title string
		Body  []byte
	}{Title: "Player", Body: bytes.Join(currentBody, []byte(", "))})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
