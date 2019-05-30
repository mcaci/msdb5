package game

import (
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/app/gamelog"
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

// Process func
func (g *Game) Process(request, origin string) []*Info {
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
	criteria := findCriteria(g, request, origin)
	actingPlayerIndex, actingPlayer, err := g.players.Find(criteria)
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
	infos := make([]*Info, 0)
	for _, pl := range g.players {
		if pl.IsSameHost(origin) {
			infos = append(infos, NewInfo(pl.Host(), gamelog.ToMe(g), err))
		}
		infos = append(infos, NewInfo(pl.Host(), gamelog.ToAll(g), err))
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
		infos = []*Info{NewInfo(fmt.Sprintf("Callers: %+v; Others: %+v", scoreTeam1, scoreTeam2), "", nil)}
	}

	return infos
}
