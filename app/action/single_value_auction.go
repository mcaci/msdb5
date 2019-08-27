package action

import (
	"strconv"

	"github.com/mcaci/ita-cards/set"

	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type auctionData struct {
	currentPlayer *player.Player
	players       team.Players
	score         *auction.Score
	update        func(auction.Score)
	side          *set.Cards
	setShowSide   func(bool, uint8)
}

func (a auctionData) valueSet(val string) {
	score, err := strconv.Atoi(val)
	toFold := player.Folded(a.currentPlayer) || err != nil || !auction.CheckScores(*a.score, auction.Score(score))
	if toFold {
		a.currentPlayer.Fold()
	}
	newScore := auction.Update(*a.score, auction.Score(score))
	a.update(newScore)
	if newScore < 120 {
		return
	}
	for _, p := range a.players {
		if p == a.currentPlayer {
			continue
		}
		p.Fold()
	}
	if len(*a.side) == 0 {
		return
	}
	a.setShowSide(len(*a.side) > 0, auction.SideCards(newScore))
}
