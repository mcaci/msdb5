package collect

import (
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/phase"
	"github.com/mcaci/msdb5/v2/dom/team"
)

// Collector func
func Collector(p phase.ID, all team.Players, side *set.Cards, played *set.Cards) (collector func() *set.Cards) {
	collector = func() *set.Cards { return &set.Cards{} }
	switch p {
	case phase.PlayingCards:
		if len(*played) == 5 {
			collector = func() *set.Cards { return played }
		}
	case phase.End:
		collector = NewAllCards(all, side, played).Set
	}
	return
}

var none *set.Cards = &set.Cards{}

type RoundCards struct {
	onTable *set.Cards
}

func NewRoundCards(onTable *set.Cards) *RoundCards { return &RoundCards{onTable: onTable} }
func (rc *RoundCards) Set() *set.Cards {
	if len(*rc.onTable) == 5 {
		return rc.onTable
	}
	return none
}

type AllCards struct {
	all           team.Players
	side, onTable *set.Cards
}

func NewAllCards(players team.Players, side, onTable *set.Cards) *AllCards {
	return &AllCards{
		all:     players,
		side:    side,
		onTable: onTable,
	}
}

func (ac *AllCards) Set() *set.Cards {
	leftoverCards := &set.Cards{}
	set.Move(ac.side, leftoverCards)
	set.Move(ac.onTable, leftoverCards)
	for _, p := range ac.all {
		set.Move(p.Hand(), leftoverCards)
	}
	return leftoverCards
}
