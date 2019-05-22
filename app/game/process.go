package game

import (
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/app"

	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

// Process func
func (g *Game) Process(request, origin string) *app.Info {
	// phase step
	currentPhase := g.CurrentPhase()
	inputPhase := phase.ToID(request)
	if currentPhase != inputPhase {
		return app.NewInfo("", "", fmt.Errorf("Phase is not %d but %d", inputPhase, currentPhase))
	}

	// find step
	playerInTurn := g.PlayerInTurn()
	findPredicate := find(g, request, origin)
	_, actingPlayer, err := g.Players().Find(findPredicate)
	if err != nil {
		return app.NewInfo("", "", err)
	}

	// do step
	if err := play(g, actingPlayer, request, origin); err != nil {
		return app.NewInfo("", "", err)
	}

	// log action to file
	toFile(currentPhase, playerInTurn, g)

	// next player step
	g.playerInTurn = nextPlayer(g, currentPhase, g.playerInTurn)

	// next phase
	g.phase = nextPhase(g, request)

	// log action to players
	info := app.NewInfo(infoForAll(currentPhase, *g), infoForMe(*playerInTurn, currentPhase, *g), err)
	g.Log(request, origin, err)

	// clean phase
	if len(*g.PlayedCards()) >= 5 {
		g.playedCards.Clear()
	}

	// process end game
	phaseAtEndTurn := g.CurrentPhase()
	if phaseAtEndTurn == phase.End {
		scorers := make([]player.Scorer, 0)
		for _, p := range g.Players() {
			scorers = append(scorers, p)
		}
		// TODO: acting player is not the caller
		scoreTeam1, scoreTeam2 := team.Score(actingPlayer, g.Companion(), scorers...)
		info = app.NewInfo(fmt.Sprintf("Callers: %+v; Others: %+v", scoreTeam1, scoreTeam2), "", nil)
	}

	return info
}
