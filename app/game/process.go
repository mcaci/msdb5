package game

import (
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/app/gamelog"
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

// Process func
func (g *Game) Process(request, origin string) {
	// phase step
	currentPhase := g.phase
	inputPhase, err := phase.ToID(request)
	if err != nil {
		gamelog.ToConsole(g, g.sender(origin), request, err)
		g.CurrentPlayer().ReplyWith(err.Error())
		return
	}
	if currentPhase != inputPhase {
		gamelog.ToConsole(g, g.sender(origin), request, err)
		g.CurrentPlayer().ReplyWith(fmt.Sprintf("Phase is not %d but %d", inputPhase, currentPhase))
		return
	}

	// find step
	criteria := findCriteria(g, request, origin)
	_, actingPlayer, err := g.players.Find(criteria)
	if err != nil {
		gamelog.ToConsole(g, g.sender(origin), request, err)
		g.sender(origin).ReplyWith(err.Error())

		return
	}
	trackActing(&g.lastPlaying, actingPlayer)

	// do step
	if err := play(g, request, origin); err != nil {
		gamelog.ToConsole(g, g.sender(origin), request, err)
		g.CurrentPlayer().ReplyWith(err.Error())
		return
	}

	// log action to file
	gamelog.Write(g)

	// next phase
	g.phase = nextPhase(g, request)

	// next player step
	nextPlayer(g, request, origin, currentPhase)

	// log action to players
	g.LastPlayer().ReplyWith(gamelog.ToMe(g))
	for _, pl := range g.players {
		pl.ReplyWith(gamelog.ToAll(g))
	}
	gamelog.ToConsole(g, g.sender(origin), request, err)

	// clean phase
	if g.cardsOnTheBoard() >= 5 {
		g.playedCards.Clear()
	}

	// process end game
	phaseAtEndTurn := g.phase
	if phaseAtEndTurn == phase.End {
		scorers := make([]team.Scorer, 0)
		for _, p := range g.players {
			scorers = append(scorers, p)
		}
		scoreTeam1, scoreTeam2 := team.Score(g.caller, g.companion, scorers...)
		for _, p := range g.players {
			p.ReplyWith(fmt.Sprintf("Callers: %+v; Others: %+v", scoreTeam1, scoreTeam2))
		}
	}
}
