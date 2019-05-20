package orchestrator

import (
	"github.com/nikiforosFreespirit/msdb5/app/action"
	"github.com/nikiforosFreespirit/msdb5/app/action/clean"
	"github.com/nikiforosFreespirit/msdb5/app/action/execute/auction"
	"github.com/nikiforosFreespirit/msdb5/app/action/execute/exchange"
	"github.com/nikiforosFreespirit/msdb5/app/action/execute/join"
	"github.com/nikiforosFreespirit/msdb5/app/action/execute/nominate"
	"github.com/nikiforosFreespirit/msdb5/app/action/execute/play"
	"github.com/nikiforosFreespirit/msdb5/app/action/find"

	"github.com/nikiforosFreespirit/msdb5/dom/deck"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

func NewFinder(requestname, request, origin string, currentPlayer *player.Player) (finder action.Finder) {
	switch requestname {
	case "Join":
		finder = find.NewJoinFinder()
	default:
		finder = find.NewPlayerFinder(origin, currentPlayer)
	}
	return
}

func NewExecuter(requestname, request, origin string, o *Orchestrator) (executer action.Executer) {
	switch requestname {
	case "Join":
		executer = join.NewJoin(request, origin)
	case "Auction":
		executer = auction.NewAuction(request, origin, o.game.AuctionScore())
	case "Exchange":
		executer = exchange.NewExchangeCards(request, origin, o.game.SideDeck())
	case "Companion":
		executer = nominate.NewCompanion(request, origin, o.game.Players(), o.game.SetCompanion)
	case "Card":
		executer = play.NewPlay(request, origin, o.game.Players(),
			o.game.PlayedCards(), o.game.SideDeck(), o.game.BriscolaSeed())
	}
	return
}

func NewCleaner(requestname string, playedCards *deck.Cards) (cleaner action.Cleaner) {
	switch requestname {
	case "Card":
		cleaner = clean.NewCleaner(playedCards)
	}
	return
}
