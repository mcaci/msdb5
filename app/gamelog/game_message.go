package gamelog

import (
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

func createInGameMsg(gameInfo selfInformer, pl *player.Player) string {
	me := fmt.Sprintf("Player: %+v\n", pl)
	if gameInfo.Phase() == phase.ExchangingCards {
		me += fmt.Sprintf("Side deck: %+v\n", *gameInfo.SideDeck())
	}
	return me
}

func createSideGameMsg(gameInfo sidedeckInformer, quantity uint8) string {
	return fmt.Sprintf("Side deck: %+v\n", (*gameInfo.SideDeck())[:quantity])
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
