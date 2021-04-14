package srv

import (
	"bytes"
	"net/http"
)

type DrawPage struct {
	Title string
	Body  []byte
}

func Draw(w http.ResponseWriter, r *http.Request) {
	currentBody = append(currentBody, []byte(cards.Top().String()))
	p := &DrawPage{Title: "Player", Body: bytes.Join(currentBody, []byte(", "))}
	err := templates.ExecuteTemplate(w, "start.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
