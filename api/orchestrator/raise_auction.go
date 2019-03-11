package orchestrator

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/auction"
	"github.com/nikiforosFreespirit/msdb5/display"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

func (g *Game) raiseAuction(request, origin string) (all []display.Info, me []display.Info, err error) {
	playerInTurn := g.playerInTurn
	info := g.raiseAuctionData(request, origin)
	return g.Info(), g.players[playerInTurn].Info(), g.playPhase(info)
}

func (g *Game) raiseAuctionData(request, origin string) phaseData {
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
	nextPhasePredicate := auctionNextPhase
	playerPredicate := func(p *player.Player) bool { return p.Folded() }
	return phaseData{phase, find, do, nextPlayerOperator, nextPhasePredicate, playerPredicate}
}

func auctionNextPhase(players playerset.Players, searchCriteria func(*player.Player) bool) bool {
	return players.Count(searchCriteria) == 4
}
