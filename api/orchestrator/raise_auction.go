package orchestrator

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/auction"
	"github.com/nikiforosFreespirit/msdb5/display"
	"github.com/nikiforosFreespirit/msdb5/player"
)

func (g *Game) raiseAuction(request, origin string) (all []display.Info, me []display.Info, err error) {
	playerInTurn := g.playerInTurn
	info := g.raiseAuctionData(request, origin)
	return g.Info(), g.players[playerInTurn].Info(), g.playPhase(info)
}

func (g *Game) raiseAuctionData(request, origin string) dataPhase {
	data := strings.Split(request, "#")
	find := func(p *player.Player) bool { return isActive(g, p, origin) }
	do := func(p *player.Player) error {
		score := data[1]
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
	return dataPhase{scoreAuction, find, do, nextPlayerSupplier, nextPhasePredicate}
}
