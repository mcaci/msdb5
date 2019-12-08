package start

import (
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type starter interface {
	Players() team.Players
	SideDeck() *set.Cards
}

func Players(pls *team.Players) {
	for i := 0; i < 5; i++ {
		pls.Add(player.New())
	}
}

func DistributeCards(g starter, withSide bool) {
	d := set.Deck()
	for i := 0; i < set.DeckSize; i++ {
		if withSide && i >= set.DeckSize-5 {
			g.SideDeck().Add(d.Top())
		} else {
			g.Players()[i%5].Hand().Add(d.Top())
		}
	}
}
