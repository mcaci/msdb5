package game

import (
	"github.com/mcaci/ita-cards/set"

	"github.com/mcaci/msdb5/app/game/track"
	"github.com/mcaci/msdb5/dom/player"
)

func makePlayers(g *Game) {
	for i := 0; i < 5; i++ {
		g.players.Add(player.New())
	}
}

func distributeCards(g *Game, withSide bool) {
	d := set.Deck()
	for i := 0; i < set.DeckSize; i++ {
		if withSide && i >= set.DeckSize-5 {
			g.side.Add(d.Top())
		} else {
			track.Player(&g.lastPlaying, g.players[i%5])
			g.CurrentPlayer().Hand().Add(d.Top())
		}
	}
}
