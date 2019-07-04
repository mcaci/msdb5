package notify

import (
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"golang.org/x/text/message"
)

func createInGameMsg(gameInfo selfInformer, pl *player.Player) string {
	printer := message.NewPrinter(gameInfo.Lang())
	me := TranslatePlayer(*pl, printer)
	if gameInfo.Phase() == phase.ExchangingCards {
		me += printer.Sprintf("Side deck: %s\n", TranslateCards(*gameInfo.SideDeck(), printer))
	}
	return me
}

func createSideGameMsg(gameInfo sidedeckInformer, quantity uint8) string {
	printer := message.NewPrinter(gameInfo.Lang())
	return printer.Sprintf("Side deck section: %s\n", TranslateCards((*gameInfo.SideDeck())[:quantity], printer))
}

func createMlMsg(gameInfo miner) (bool, string) {
	var msg string
	canLog := true
	switch gameInfo.Phase() {
	case phase.ChoosingCompanion:
		msg = fmt.Sprintf("%s, %s, %d\n", gameInfo.CurrentPlayer().Name(), gameInfo.Companion().Name(), *(gameInfo.AuctionScore()))
	case phase.PlayingCards:
		msg = fmt.Sprintf("%s, %d\n", gameInfo.CurrentPlayer().Name(), gameInfo.LastCardPlayed())
	case phase.End:
		msg = fmt.Sprintf("%s\n", gameInfo.CurrentPlayer().Name())
	default:
		canLog = false
	}
	return canLog, msg
}
