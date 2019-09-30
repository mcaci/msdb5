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
	setShowSide   func(uint8)
	setCaller     func(player.Predicate)
}

func (a auctionData) valueSet(val string) {
	score, err := strconv.Atoi(val)
	toFold := player.Folded(a.currentPlayer) || err != nil || !auction.CheckScores(*a.score, auction.Score(score))
	if toFold {
		a.currentPlayer.Fold()
	}
	newScore := auction.Update(*a.score, auction.Score(score))
	a.update(newScore)

	if len(*a.side) > 0 {
		a.setShowSide(sideCards(uint8(newScore)))
	}

	if newScore >= 120 {
		for _, p := range a.players {
			if p == a.currentPlayer {
				continue
			}
			p.Fold()
		}
	}

	notFolded := func(p *player.Player) bool { return !player.Folded(p) }
	if team.Count(a.players, notFolded) == 1 {
		a.setCaller(notFolded)
	}
}

func sideCards(score uint8) uint8 {
	return score/90 + score/100 + score/110 + score/120 + score/120
}
