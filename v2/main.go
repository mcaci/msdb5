package main

import (
	"flag"
	"log"

	"github.com/mcaci/msdb5/v2/app/game"
	"github.com/mcaci/msdb5/v2/frw"
)

func main() {
	noSide := flag.Bool("no-side", false, "Add flag to specify no side deck is to be used.")
	flag.Parse()
	g := game.New()
	game.Setup(g, *noSide)
	game.WaitForPlayers(g, frw.WithAINames)
	game.Start(g)

	log.Println("Match over", g)
	log.Println("Score", game.Score(g))
}
