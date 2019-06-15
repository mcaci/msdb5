package gamelog

import (
	"log"
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

// ToFile func
func ToFile(gameInfo miner) {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	logger := log.New(f, "", log.LstdFlags)
	switch gameInfo.Phase() {
	case phase.ChoosingCompanion:
		logger.Printf("%s, %s, %d\n", gameInfo.CurrentPlayer().Name(), gameInfo.Companion().Name(), gameInfo.AuctionScore())
	case phase.PlayingCards:
		logger.Printf("%s, %d\n", gameInfo.CurrentPlayer().Name(), gameInfo.LastCardPlayed())
	case phase.End:
		logger.Printf("%s\n", gameInfo.CurrentPlayer().Name())
	}
}
