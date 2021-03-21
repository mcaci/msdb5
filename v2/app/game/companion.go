package game

import (
	"math/rand"
	"time"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/v2/dom/phase"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func runCompanion(g *Game) {
	for g.phase == phase.ChoosingCompanion {
		rand.Seed(time.Now().Unix())
		c := card.MustID(uint8(rand.Intn(40) + 1))
		idx, err := g.players.Index(player.IsCardInHand(*c))
		if err != nil {
			continue
		}
		g.briscolaCard = *c

		pl := g.players.At(idx)
		g.companion = pl

		// next phase
		g.phase++

		// next player: no change
	}
}

func runCompanion_v2(g struct {
	players team.Players
}, listenForId func(chan<- uint8)) struct {
	briscolaCard *card.Item
	companion    *player.Player
} {
	id := make(chan uint8)
	defer close(id)

	for {
		go listenForId(id)
		c := card.MustID(<-id)
		idx, err := g.players.Index(player.IsCardInHand(*c))
		if err != nil {
			continue
		}
		pl := g.players.At(idx)

		return struct {
			briscolaCard *card.Item
			companion    *player.Player
		}{
			briscolaCard: c,
			companion:    pl,
		}
	}
}
