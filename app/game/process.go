package game

import (
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/app/gamelog"
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

// Process func
func (g *Game) Process(request, origin string) {
	// find step
	criteria := findCriteria(g, request, origin)
	actingPlayerIndex, actingPlayer, err := g.players.Find(criteria)
	if err != nil {
		gamelog.ToConsole(g, request, err)
		criteria := findCriteria(g, "Origin", origin)
		_, originPlayer, stillErr := g.players.Find(criteria)
		if stillErr == nil {
			originPlayer.ReplyWith(err.Error())
		}
		return
	}

	// phase step
	currentPhase := g.phase
	inputPhase, err := phase.ToID(request)
	if err != nil {
		gamelog.ToConsole(g, request, err)
		actingPlayer.ReplyWith(err.Error())
		return
	}
	if currentPhase != inputPhase {
		gamelog.ToConsole(g, request, err)
		actingPlayer.ReplyWith(fmt.Sprintf("Phase is not %d but %d", inputPhase, currentPhase))
		return
	}

	// do step
	if err := play(g, actingPlayer, request, origin); err != nil {
		gamelog.ToConsole(g, request, err)
		actingPlayer.ReplyWith(err.Error())
		return
	}

	// log action to file
	gamelog.Write(g)

	// next phase
	g.phase = nextPhase(g, request)

	// next player step
	nextPlayer(g, currentPhase, actingPlayerIndex)

	// log action to players
	g.LastPlayer().ReplyWith(gamelog.ToMe(g))
	for _, pl := range g.players {
		pl.ReplyWith(gamelog.ToAll(g))
	}
	gamelog.ToConsole(g, request, err)

	// clean phase
	if g.cardsOnTheBoard() >= 5 {
		g.playedCards.Clear()
	}

	// process end game
	phaseAtEndTurn := g.phase
	if phaseAtEndTurn == phase.End {
		scorers := make([]player.Scorer, 0)
		for _, p := range g.players {
			scorers = append(scorers, p)
		}
		scoreTeam1, scoreTeam2 := team.Score(g.caller, g.companion, scorers...)
		actingPlayer.ReplyWith(fmt.Sprintf("Callers: %+v; Others: %+v", scoreTeam1, scoreTeam2))
	}
}
