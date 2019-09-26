package start

import (
	"container/list"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/app/game/track"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

func Players(pls *team.Players) {
	for i := 0; i < 5; i++ {
		pls.Add(player.New())
	}
}

func Distribute(g interface {
	CurrentPlayer() *player.Player
	LastPlaying() *list.List
	Players() team.Players
	SideDeck() *set.Cards
}, withSide bool) {
	d := set.Deck()
	for i := 0; i < set.DeckSize; i++ {
		if withSide && i >= set.DeckSize-5 {
			g.SideDeck().Add(d.Top())
		} else {
			track.Player(g.LastPlaying(), g.Players()[i%5])
			g.CurrentPlayer().Hand().Add(d.Top())
		}
	}
}
