package orchestrator

import (
	"github.com/nikiforosFreespirit/msdb5/api/action"
	"github.com/nikiforosFreespirit/msdb5/player"
)

func NewFinder(requestname, request, origin string, currentPlayer *player.Player) (finder action.Finder) {
	switch requestname {
	case "Join":
		finder = action.NewJoinFinder(request, origin)
	default:
		finder = action.NewPlayerFinder(origin, currentPlayer)
	}
	return
}

func NewExecuter(requestname, request, origin string, o *Orchestrator) (executer action.Executer) {
	switch requestname {
	case "Join":
		executer = action.NewJoin(request, origin)
	case "Auction":
		executer = action.NewAuction(request, origin, o.game.Players(), o.game.Board())
	case "Exchange":
		executer = action.NewExchangeCards(request, origin, o.game.Board().SideDeck())
	case "Companion":
		executer = action.NewCompanion(request, origin, o.game.Players(), o.game.SetCompanion)
	case "Card":
		executer = action.NewPlay(request, origin, o.game.Players(),
			o.game.Board().PlayedCards(), o.game.Board().SideDeck(), o.game.BriscolaSeed())
	}
	return
}

func NewSelector(requestname, request, origin string, o *Orchestrator) (selector action.NextPlayerSelector) {
	switch requestname {
	case "Join":
		selector = action.NewPlayerSelector()
	case "Auction":
		selector = action.NewAuction(request, origin, o.game.Players(), o.game.Board())
	case "Exchange", "Companion":
		selector = action.NewSelfPlayerSelector()
	case "Card":
		selector = action.NewPlay(request, origin, o.game.Players(),
			o.game.Board().PlayedCards(), o.game.Board().SideDeck(), o.game.BriscolaSeed())
	}
	return
}

func NewChanger(requestname, request, origin string, o *Orchestrator) (selector action.NextPhaseChanger) {
	switch requestname {
	case "Join":
		selector = action.NewPhaseChanger(o.game.Players())
	case "Auction":
		selector = action.NewAuction(request, origin, o.game.Players(), o.game.Board())
	case "Exchange":
		selector = action.NewExchangeCards(request, origin, o.game.Board().SideDeck())
	case "Companion":
		selector = action.NewCompanion(request, origin, o.game.Players(),
			o.game.SetCompanion)
	case "Card":
		selector = action.NewPlay(request, origin, o.game.Players(),
			o.game.Board().PlayedCards(), o.game.Board().SideDeck(), o.game.BriscolaSeed())
	}
	return
}
