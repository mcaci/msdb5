package game

import (
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/app/gamelog"
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

// Process func
func (g *Game) Process(request, origin string) *Info {
	// phase step
	currentPhase := g.phase
	inputPhase, err := phase.ToID(request)
	if err != nil {
		gamelog.ToConsole(g, request, err)
		return NewErrorInfo(err)
	}
	if currentPhase != inputPhase {
		gamelog.ToConsole(g, request, err)
		return NewErrorInfo(fmt.Errorf("Phase is not %d but %d", inputPhase, currentPhase))
	}

	// find step
	findPredicate := find(g, request, origin)
	actingPlayerIndex, actingPlayer, err := g.playersRef().Find(findPredicate)
	if err != nil {
		gamelog.ToConsole(g, request, err)
		return NewErrorInfo(err)
	}

	// do step
	if err := play(g, actingPlayer, request, origin); err != nil {
		gamelog.ToConsole(g, request, err)
		return NewErrorInfo(err)
	}

	// log action to file
	gamelog.Write(g)

	// next phase
	g.phase = nextPhase(g, request)

	// next player step
	nextPlayer(g, currentPhase, actingPlayerIndex)

	// log action to players
	info := NewInfo(gamelog.ToAll(g), gamelog.ToMe(g), err)
	gamelog.ToConsole(g, request, err)

	// clean phase
	if g.cardsOnTheBoard() >= 5 {
		g.playedCards.Clear()
	}

	// process end game
	phaseAtEndTurn := g.phase
	if phaseAtEndTurn == phase.End {
		scorers := make([]player.Scorer, 0)
		for _, p := range g.playersRef() {
			scorers = append(scorers, p)
		}
		scoreTeam1, scoreTeam2 := team.Score(g.caller, g.companion, scorers...)
		info = NewInfo(fmt.Sprintf("Callers: %+v; Others: %+v", scoreTeam1, scoreTeam2), "", nil)
	}

	return info
}
