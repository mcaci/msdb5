package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/mcaci/msdb5/v2/app/game"
	"github.com/mcaci/msdb5/v2/frw/srv"
)

func main() {
	noSide := flag.Bool("no-side", false, "Flag to specify no side deck is to be used. Default: false (no side deck used)")
	aiGame := flag.Bool("ai-game", false, "Flag to specify if ai plays the game (for tests). Default: false (webapp at :8080)")
	flag.Parse()

	if !*aiGame {
		http.HandleFunc("/", srv.Home)
		http.HandleFunc("/start/", srv.Start)
		http.HandleFunc("/draw/", srv.Draw)
		log.Fatal(http.ListenAndServe(":8080", nil))
	}

	// setup game
	g := game.NewGame(&game.Options{
		WithSide: !*noSide,
	})

	game.Start(g)

	log.Println("Match over", g)
	log.Println("Score", game.ScoreGrpc(g))
}
