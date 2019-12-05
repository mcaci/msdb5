package start

import (
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

func Players(pls *team.Players) {
	for i := 0; i < 5; i++ {
		pls.Add(player.New())
	}
}

func DistributeCards(g interface{ Players() team.Players }, withSide bool) set.Cards {
	d := set.Deck()
	for i := 0; i < set.DeckSize; i++ {
		if withSide && i >= set.DeckSize-5 {
			break
		} else {
			g.Players()[i%5].Hand().Add(d.Top())
		}
	}
	return d
}
