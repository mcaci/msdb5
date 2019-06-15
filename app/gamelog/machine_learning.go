package gamelog

import (
	"io"
	"os"

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
	SideDeck() deck.Cards
}

// OpenFile func
func OpenFile() (*os.File, error) {
	return os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
}

// ToFile func
func ToFile(gameInfo miner, writer io.Writer) {
	canLog, msg := createMlMsg(gameInfo)
	if !canLog {
		return
	}
	write(writer, msg)
}
