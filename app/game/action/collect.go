package action

import (
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/phase"
	"github.com/mcaci/msdb5/dom/team"
)

// Collector func
func Collector(g interface{ Phase() phase.ID }, all team.Players, side *set.Cards, played *set.Cards) (collector func() *set.Cards) {
	collector = func() *set.Cards { return &set.Cards{} }
	switch g.Phase() {
	case phase.PlayingCards:
		if len(*played) == 5 {
			collector = func() *set.Cards { return played }
		}
	case phase.End:
		collector = allCards{all, side, played}.leftInGame
	}
	return
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
