package srv

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mcaci/msdb5/v2/app/game"
)

type StartPage struct {
	Title string
	Body  []byte
	Msg   []byte
}

func Start(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	playername := r.Form["playername"][0]
	gamename := r.Form["gamename"][0]
	v := &StartPage{Title: fmt.Sprintf("Welcome %s! Game %s has started.", playername, gamename)}
	switch r.Form["type"][0] {
	case "create":
		if g != nil {
			log.Printf("Game with gamename %q already exists, cannot creata game %q too", n, gamename)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		g = game.NewGame(&game.Options{WithSide: true})
		n = gamename
		v.Msg = []byte(fmt.Sprintf("new game created with gamename %q", gamename))
		log.Printf("Game created with gamename %q", gamename)
	case "join":
		if g == nil {
			log.Println("No games existing yet", gamename)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		if gamename != n {
			log.Printf("game %s not found", gamename)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		v.Msg = []byte(fmt.Sprintf("joining game %q", gamename))
	default:
		log.Printf("unknown %q option", r.Form["type"][0])
		http.Error(w, "Did not understand the action", http.StatusInternalServerError)
		return
	}
	log.Println(r.Form)
	err = templates.ExecuteTemplate(w, "start.html", v)
	if err != nil {
		http.NotFound(w, r)
		return
	}
}
