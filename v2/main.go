package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/mcaci/msdb5/v2/app/game"
	"github.com/mcaci/msdb5/v2/frw"
)

func main() {
	runGame()
}

func runGame() {
	noSide := flag.Bool("no-side", false, "Add flag to specify no side deck is to be used.")
	flag.Parse()
	g := game.New()
	game.Setup(g, *noSide)
	game.WaitForPlayers(g, frw.WithAINames)
	game.Start(g)

	log.Println("Match over", g)
	log.Println("Score", game.Score(g))
}

func welcome(rw http.ResponseWriter, r *http.Request) {
	log.Print("New Request: ", r)
	rw.Write([]byte("Welcome guest"))
}
