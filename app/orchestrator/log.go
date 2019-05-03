package orchestrator

import (
	"fmt"
	"log"
	"os"

	"github.com/nikiforosFreespirit/msdb5/dom/player"

	"github.com/nikiforosFreespirit/msdb5/app/action"
	"github.com/nikiforosFreespirit/msdb5/app/action/execute/nominate"
	"github.com/nikiforosFreespirit/msdb5/app/action/execute/play"
	"github.com/nikiforosFreespirit/msdb5/app/game"
)

func toFile(actionExec action.Executer, p *player.Player, g *game.Game) {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	logger := log.New(f, "", log.LstdFlags)
	switch actionExec.(type) {
	case *nominate.CompanionStruct:
		logger.Printf("%s, %s, %d\n", p.Name(), g.Companion().Name(), g.Board().AuctionScore())
	case *play.PlayCardStruct:
		idx := len(*g.Board().PlayedCards()) - 1
		logger.Printf("%s, %d\n", p.Name(), (*g.Board().PlayedCards())[idx])
	}
}

func infoForAll(currentPhase game.Phase, gameInfo game.Game) string {
	all := fmt.Sprintf("Game: %+v", gameInfo)
	board := gameInfo.Board()
	isSideDeckUsed := len(*board.SideDeck()) > 0
	if currentPhase == game.InsideAuction && isSideDeckUsed {
		if board.AuctionScore() >= 90 {
			all += fmt.Sprintf("First card: %+v", (*board.SideDeck())[0])
		}
		if board.AuctionScore() >= 100 {
			all += fmt.Sprintf("Second card: %+v", (*board.SideDeck())[1])
		}
		if board.AuctionScore() >= 110 {
			all += fmt.Sprintf("Third card: %+v", (*board.SideDeck())[2])
		}
		if board.AuctionScore() >= 120 {
			all += fmt.Sprintf("Fourth card: %+v", (*board.SideDeck())[3])
			all += fmt.Sprintf("Fifth card: %+v", (*board.SideDeck())[4])
		}
	}
	return all
}

func infoForMe(currentPlayer player.Player, currentPhase game.Phase, gameInfo game.Game) string {
	me := fmt.Sprintf("%+v", currentPlayer)
	if currentPhase == game.ExchangingCards {
		me += fmt.Sprintf("Side deck: %+v", gameInfo.Board().SideDeck())
	}
	return me
}
