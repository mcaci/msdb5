package pl

import (
	"io"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/dom/player"
	"golang.org/x/text/message"
)

type lastPlInformer interface {
	Briscola() card.Item
	CurrentPlayer() *player.Player
	LastPlayer() *player.Player
}

func ToLastPl(g lastPlInformer, printer *message.Printer) {
	if g.LastPlayer() != g.CurrentPlayer() {
		io.WriteString(g.LastPlayer(), translatePlayer(g.LastPlayer(), g.Briscola(), printer))
	}
}
