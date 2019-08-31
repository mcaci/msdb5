package msg

import (
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/player"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type selfInformer interface {
	Phase() phase.ID
	SideDeck() *set.Cards
}

// CreateInGameMsg func
func CreateInGameMsg(gameInfo selfInformer, pl *player.Player, l language.Tag) string {
	printer := message.NewPrinter(l)
	me := printer.Sprintf("Player: (Name: %s, Cards: %+v, Pile: %+v, Has folded? %t)",
		pl.Name(), TranslateCards(*pl.Hand(), printer), TranslateCards(*pl.Pile(), printer), player.Folded(pl))
	if gameInfo.Phase() == phase.ExchangingCards {
		me += " " + printer.Sprintf("Side deck: %s\n", TranslateCards(*gameInfo.SideDeck(), printer))
	}
	return me
}
