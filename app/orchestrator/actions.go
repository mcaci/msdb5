package orchestrator

import (
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/dom/deck"

	"github.com/nikiforosFreespirit/msdb5/app/action"
	"github.com/nikiforosFreespirit/msdb5/app/action/clean"
	"github.com/nikiforosFreespirit/msdb5/app/action/execute/auction"
	"github.com/nikiforosFreespirit/msdb5/app/action/execute/exchange"
	"github.com/nikiforosFreespirit/msdb5/app/action/execute/join"
	"github.com/nikiforosFreespirit/msdb5/app/action/execute/nominate"
	"github.com/nikiforosFreespirit/msdb5/app/action/execute/play"
	"github.com/nikiforosFreespirit/msdb5/app/action/find"

	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/playerset"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
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
		executer = auction.NewAuction(request, origin, o.game.Board().AuctionScore())
	case "Exchange":
		executer = exchange.NewExchangeCards(request, origin, o.game.Board().SideDeck())
	case "Companion":
		executer = nominate.NewCompanion(request, origin, o.game.Players(), o.game.SetCompanion)
	case "Card":
		executer = play.NewPlay(request, origin, o.game.Players(),
			o.game.Board().PlayedCards(), o.game.Board().SideDeck(), o.game.BriscolaSeed())
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

func endGame(players playerset.Players, companion player.ScoreCounter) (string, string, error) {
	caller, _ := players.Find(func(p *player.Player) bool { return p.NotFolded() })
	team1, team2 := new(team.BriscolaTeam), new(team.BriscolaTeam)
	team1.Add(caller, companion)
	for _, pl := range players {
		if pl != caller && pl != companion {
			team2.Add(pl)
		}
	}
	return fmt.Sprintf("Callers: %+v; Others: %+v", team1.Score(), team2.Score()), "", nil
}
