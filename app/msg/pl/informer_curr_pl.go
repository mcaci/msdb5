package pl

import (
	"fmt"
	"io"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/phase"
	"github.com/mcaci/msdb5/dom/player"
	"golang.org/x/text/message"
)

type currPlInformer interface {
	Briscola() card.Item
	CurrentPlayer() *player.Player
	Phase() phase.ID
	SideDeck() *set.Cards
}

func ToNewPl(g currPlInformer, printer *message.Printer) {
	if g.Phase() == phase.ExchangingCards {
		io.WriteString(g.CurrentPlayer(), fmt.Sprintf("%s: %s\n", sideDeckRef(printer), translateCards(*g.SideDeck(), printer)))
	}
	io.WriteString(g.CurrentPlayer(), translatePlayer(g.CurrentPlayer(), g.Briscola(), printer))
}
