package msg

import (
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/auction"
	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/deck"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

type miner interface {
	AuctionScore() *auction.Score
	Companion() *player.Player
	CurrentPlayer() *player.Player
	LastCardPlayed() card.ID
	Phase() phase.ID
	// not registerd yet
	IsSideUsed() bool
	SideDeck() *deck.Cards
}

// CreateMlMsg func
func CreateMlMsg(gameInfo miner) (bool, string) {
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
