package orchestrator

import (
	"github.com/nikiforosFreespirit/msdb5/auction"
	"github.com/nikiforosFreespirit/msdb5/player"
)

// RaiseAuction func
func (g *Game) RaiseAuction(score, origin string) (err error) {
	find := func(p *player.Player) bool { return isActive(g, p, origin) }
	do := func(p *player.Player) error {
		auction.CheckAndUpdate(score, p.Folded, p.Fold, g.info.AuctionScore, g.info.SetAuctionScore)
		return nil
	}
	nextPlayerSupplier := func() uint8 {
		winnerIndex := (g.playerInTurn + 1) % 5
		for g.players[winnerIndex].Folded() {
			winnerIndex = (winnerIndex + 1) % 5
		}
		return winnerIndex
	}
	nextPhasePredicate := func() bool { return g.players.Count(folded) == 4 }
	return g.playPhase(scoreAuction, find, do, nextPlayerSupplier, nextPhasePredicate)
}
