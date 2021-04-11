package main

import (
	"flag"
	"log"

	"github.com/mcaci/msdb5/v2/app/game"
)

func main() {
	noSide := flag.Bool("no-side", false, "Flag to specify no side deck is to be used. Default: false (no side deck used)")
	flag.Parse()

	// setup game
	g := game.NewGame(&game.Options{
		WithSide: !*noSide,
	})

	game.WaitForPlayers(g)
	game.Start(g)

	log.Println("Match over", g)
	log.Println("Score", game.ScoreGrpc(g))
}
