package frw

import (
	"log"

	"github.com/mcaci/msdb5/v2/app/game"
	"github.com/mcaci/msdb5/v2/app/listen"
)

func Game(noSide bool) {
	g := game.New()
	game.Setup(g, noSide)
	game.WaitForPlayers(g, listen.WithAINames)
	game.Start(g)

	log.Println("Match over", g)
	log.Println("Score", game.Score(g))
}
