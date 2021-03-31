package companion

import (
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func Run(players team.Players, listenForId func(chan<- uint8)) struct {
	Briscola  briscola.Card
	Companion uint8
} {
	id := make(chan uint8)
	defer close(id)
	go listenForId(id)
	return Round(*briscola.MustID(<-id), players)
}
