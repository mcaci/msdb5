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
	phase := scoreAuction
	find := func(p *player.Player) bool { return isExpectedPlayer(p, g, origin) }
	do := func(p *player.Player) error {
		score := data[1]
		auction.CheckAndUpdate(score, p.Folded, p.Fold, g.info.AuctionScore, g.info.SetAuctionScore)
		return nil
	}
	nextPlayerOperator := func(playerInTurn uint8) uint8 {
		winnerIndex := nextPlayer(playerInTurn)
		for g.players[winnerIndex].Folded() {
			winnerIndex = nextPlayer(winnerIndex)
		}
		return winnerIndex
	}
	nextPhasePredicate := func() bool { return g.players.Count(folded) == 4 }
	return dataPhase{phase, find, do, nextPlayerOperator, nextPhasePredicate}
}
