package collect

import (
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/team"
)

type RoundCards struct {
	onTable *set.Cards
}

func NewRoundCards(onTable *set.Cards) *RoundCards { return &RoundCards{onTable: onTable} }
func (rc *RoundCards) Set() *set.Cards {
	if len(*rc.onTable) == 5 {
		return rc.onTable
	}
	return &set.Cards{}
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
