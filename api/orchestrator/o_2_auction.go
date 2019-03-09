package orchestrator

import (
	"github.com/nikiforosFreespirit/msdb5/auction"
	"github.com/nikiforosFreespirit/msdb5/player"
)

// RaiseAuction func
func (g *Game) RaiseAuction(score, origin string) (err error) {
	find := func(p *player.Player) bool { return isActive(g, p, origin) }
	nextPlayerSupplier := func() uint8 {
		winnerIndex := (g.playerInTurn + 1) % 5
		for g.players[winnerIndex].Folded() {
			winnerIndex = (winnerIndex + 1) % 5
		}
		return winnerIndex
	}
	nextPhasePredicate := func() bool { return g.players.Count(folded) == 4 }
	return g.raiseAuction(scoreAuction, score, origin, find, nextPlayerSupplier, nextPhasePredicate)
}

func (g *Game) raiseAuction(phase phase, score, origin string, find func(*player.Player) bool, nextPlayerSupplier func() uint8, nextPhasePredicate func() bool) (err error) {
	if err = g.phaseCheck(phase); err != nil {
		return
	}
	p, err := g.players.Find(find)
	if err != nil {
		return
	}
	auction.CheckAndUpdate(score, p.Folded, p.Fold, g.info.AuctionScore, g.info.SetAuctionScore)
	g.nextPlayer(nextPlayerSupplier)
	g.nextPhase(nextPhasePredicate)
	return
}
