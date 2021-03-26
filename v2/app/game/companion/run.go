package companion

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func Run(players team.Players, listenForId func(chan<- uint8)) struct {
	Briscola  *card.Item
	Companion *player.Player
} {
	id := make(chan uint8)
	defer close(id)

	for {
		go listenForId(id)
		c := card.MustID(<-id)
		return Round(c, players)
	}
}
