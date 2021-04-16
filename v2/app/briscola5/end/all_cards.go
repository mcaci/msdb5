package end

import (
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
	"github.com/mcaci/msdb5/v2/dom/team"
)

type allCards struct {
	all     team.Players
	side    briscola5.Side
	onTable *briscola5.PlayedCards
}

func newAllCards(players team.Players, side briscola5.Side, onTable *briscola5.PlayedCards) *allCards {
	return &allCards{
		all:     players,
		side:    side,
		onTable: onTable,
	}
}

func (ac *allCards) Pile() *set.Cards {
	leftoverCards := &set.Cards{}
	set.Move(&ac.side.Cards, leftoverCards)
	set.Move(ac.onTable.Cards, leftoverCards)
	for _, p := range ac.all {
		set.Move(p.Hand(), leftoverCards)
	}
	return leftoverCards
}
