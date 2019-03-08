package orchestrator

import (
	"log"

	"github.com/nikiforosFreespirit/msdb5/briscola"
	"github.com/nikiforosFreespirit/msdb5/player"
)

func winner(g *Game) uint8 {
	return (g.playerInTurn + briscola.IndexOfWinningCard(*g.info.PlayedCards(), g.companion.Card().Seed()) + 1) % 5
}

func logEndRound(g *Game, request, origin string, err error) {
	playerLogged, _ := g.players.Find(func(p *player.Player) bool { return isRemoteHost(p, origin) })
	log.Printf("New Action by %s\n", playerLogged.Name().Display())
	log.Printf("Action is %s\n", request)
	log.Printf("Any error raised: %v\n", err)
	log.Printf("Game info after action: %s\n", g.String())
}
