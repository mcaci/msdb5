package game

import (
	"fmt"
	"log"
	"os"

	"github.com/nikiforosFreespirit/msdb5/dom/player"

	"github.com/nikiforosFreespirit/msdb5/app/phase"
)

func toFile(current phase.ID, p *player.Player, g *Game) {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	logger := log.New(f, "", log.LstdFlags)
	switch current {
	case phase.ChosingCompanion:
		logger.Printf("%s, %s, %d\n", p.Name(), g.Companion().Name(), g.AuctionScore())
	case phase.PlayingCards:
		idx := len(*g.PlayedCards()) - 1
		logger.Printf("%s, %d\n", p.Name(), (*g.PlayedCards())[idx])
	}
}

func infoForAll(currentPhase phase.ID, gameInfo Game) string {
	all := fmt.Sprintf("Game: %+v", gameInfo)
	sideDeck := gameInfo.SideDeck()
	isSideDeckUsed := len((*sideDeck)) > 0
	if currentPhase == phase.InsideAuction && isSideDeckUsed {
		score := gameInfo.AuctionScore()
		if *score >= 90 {
			all += fmt.Sprintf("First card: %+v", (*sideDeck)[0])
		}
		if *score >= 100 {
			all += fmt.Sprintf("Second card: %+v", (*sideDeck)[1])
		}
		if *score >= 110 {
			all += fmt.Sprintf("Third card: %+v", (*sideDeck)[2])
		}
		if *score >= 120 {
			all += fmt.Sprintf("Fourth card: %+v", (*sideDeck)[3])
			all += fmt.Sprintf("Fifth card: %+v", (*sideDeck)[4])
		}
	}
	return all
}

func infoForMe(currentPlayer player.Player, currentPhase phase.ID, gameInfo Game) string {
	me := fmt.Sprintf("%+v", currentPlayer)
	if currentPhase == phase.ExchangingCards {
		me += fmt.Sprintf("Side deck: %+v", gameInfo.SideDeck())
	}
	return me
}
