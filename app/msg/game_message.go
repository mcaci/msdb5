package msg

import (
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/deck"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type selfInformer interface {
	Lang() language.Tag
	Phase() phase.ID
	SideDeck() *deck.Cards
}

// CreateInGameMsg func
func CreateInGameMsg(gameInfo selfInformer, pl *player.Player) string {
	printer := message.NewPrinter(gameInfo.Lang())
	me := printer.Sprintf("(Name: %s, Cards: %+v, Pile: %+v, Has folded? %t)",
		pl.Name(), TranslateCards(*pl.Hand(), printer), TranslateCards(*pl.Pile(), printer), player.Folded(pl))
	if gameInfo.Phase() == phase.ExchangingCards {
		me += printer.Sprintf("Side deck: %s\n", TranslateCards(*gameInfo.SideDeck(), printer))
	}
	return me
}
