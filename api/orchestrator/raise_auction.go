package orchestrator

import (
	"errors"
	"strings"

	"github.com/nikiforosFreespirit/msdb5/auction"
	"github.com/nikiforosFreespirit/msdb5/display"
	"github.com/nikiforosFreespirit/msdb5/player"
)

func (g *Game) raiseAuction(request, origin string) (all []display.Info, me []display.Info, err error) {
	data := strings.Split(request, "#")
	action := data[0]
	score := data[1]
	playerInTurn := g.playerInTurn
	if action == "Auction" {
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
		err = g.playPhase(scoreAuction, find, do, nextPlayerSupplier, nextPhasePredicate)
		return g.Info(), g.players[playerInTurn].Info(), err
	}
	return g.Info(), g.players[playerInTurn].Info(), errors.New("AUCTION action not invoked")
}
