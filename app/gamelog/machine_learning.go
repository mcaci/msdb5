package gamelog

import (
	"log"
	"os"

	"github.com/nikiforosFreespirit/msdb5/app/phase"
)

// Write func
func Write(gameInfo informer) {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	logger := log.New(f, "", log.LstdFlags)
	switch gameInfo.Phase() {
	case phase.ChosingCompanion:
		logger.Printf("%s, %s, %d\n", gameInfo.CurrentPlayer().Name(), gameInfo.Companion().Name(), gameInfo.AuctionScore())
	case phase.PlayingCards:
		logger.Printf("%s, %d\n", gameInfo.CurrentPlayer().Name(), gameInfo.LastCardPlayed())
	}
}
