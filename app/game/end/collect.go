package end

import (
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/phase"
	"github.com/mcaci/msdb5/dom/team"
)

// Collector func
func Collector(ph phase.ID, all team.Players, side *set.Cards, played *set.Cards) func() *set.Cards {
	if ph == phase.PlayingCards && len(*played) == 5 {
		return func() *set.Cards { return played }
	}
	if ph == phase.End {
		return allCards{all, side, played}.leftInGame
	}
	return func() *set.Cards { return &set.Cards{} }
}

type allCards struct {
	all           team.Players
	side, onTable *set.Cards
}

func (a allCards) leftInGame() *set.Cards {
	leftoverCards := &set.Cards{}
	set.Move(a.side, leftoverCards)
	set.Move(a.onTable, leftoverCards)
	for _, p := range a.all {
		set.Move(p.Hand(), leftoverCards)
	}
	return leftoverCards
}
