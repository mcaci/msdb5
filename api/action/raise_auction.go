package action

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/auction"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

// RaiseAuctionData func
func RaiseAuctionData(g *game.Game, request, origin string) Data {
	data := strings.Split(request, "#")
	phase := game.ScoreAuction
	find := func(p *player.Player) bool { return isExpectedPlayer(p, g.PlayerInTurn(), origin) }
	do := func(p *player.Player) error {
		score := data[1]
		auction.CheckAndUpdate(score, p.Folded, p.Fold, g.Board().AuctionScore, g.Board().SetAuctionScore)
		return nil
	}
	nextPlayerOperator := func(playerInTurn uint8) uint8 {
		winnerIndex := nextPlayer(playerInTurn)
		for g.Players()[winnerIndex].Folded() {
			winnerIndex = nextPlayer(winnerIndex)
		}
		return winnerIndex
	}
	nextPhasePredicate := auctionNextPhase
	playerPredicate := func(p *player.Player) bool { return p.Folded() }
	return Data{phase, find, do, nextPlayerOperator, nextPhasePredicate, playerPredicate}
}

func auctionNextPhase(players playerset.Players, searchCriteria func(*player.Player) bool) bool {
	return players.Count(searchCriteria) == 4
}
