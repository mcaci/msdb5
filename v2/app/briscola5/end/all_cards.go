package end

import (
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/misc"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
)

type allCards struct {
	all     misc.Players
	side    briscola5.Side
	onTable *briscola.PlayedCards
}

func newAllCards(players misc.Players, side briscola5.Side, onTable *briscola.PlayedCards) *allCards {
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
