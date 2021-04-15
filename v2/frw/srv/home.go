package srv

import (
	"html/template"
	"net/http"
)

var home = template.Must(template.ParseFiles("assets/home.html"))

func Home(w http.ResponseWriter, r *http.Request) {
	err := home.Execute(w, &struct {
		Title string
		Body  []byte
		Msg   []byte
	}{Title: "Home page", Body: []byte("default")})
	if err != nil {
		http.NotFound(w, r)
		return
	}
}
