package srv

import (
	"net/http"
)

type HomePage struct {
	Title string
	Body  []byte
	Msg   []byte
}

func Home(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "home.html", &HomePage{Title: "Home page", Body: []byte("default")})
	if err != nil {
		http.NotFound(w, r)
		return
	}
}
