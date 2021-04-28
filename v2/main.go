package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/mcaci/msdb5/v2/app/briscola5"
	"github.com/mcaci/msdb5/v2/frw/srv"
)

func main() {
	noSide := flag.Bool("no-side", false, "Flag to specify no side deck is to be used. Default: false (no side deck used)")
	aiGame := flag.Bool("ai-game", false, "Flag to specify if ai plays the game (for tests). Default: false (webapp at :8080)")
	flag.Parse()

	if !*aiGame {
		http.HandleFunc("/", srv.Home)
		http.HandleFunc("/start/", srv.Start)
		http.HandleFunc("/play/", srv.Play)
		http.HandleFunc("/refresh/", srv.Refresh)
		log.Fatal(http.ListenAndServe(":8080", nil))
	}

	// setup game
	g := briscola5.NewGame(&briscola5.Options{
		WithSide: !*noSide,
	})

	briscola5.Start(g)

	log.Println("Match over", g)
	log.Println("Score", briscola5.ScoreGrpc(g))
}
