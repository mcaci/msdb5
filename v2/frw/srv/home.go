package srv

import (
	"log"
	"net/http"
)

type HomePage struct {
	Title string
	Body  []byte
	Msg   []byte
}

func Home(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "home.html", &HomePage{Title: "Start", Body: []byte("test")})
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}
}
