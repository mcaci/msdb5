package srv

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "home.html", &Page{Title: "Start", Body: []byte("test")})
	if err != nil {
		http.NotFound(w, r)
		return
	}
}
