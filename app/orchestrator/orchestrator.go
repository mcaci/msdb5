package orchestrator

import (
	"fmt"
	"strings"

	"github.com/nikiforosFreespirit/msdb5/app"
	"github.com/nikiforosFreespirit/msdb5/app/action/clean"
	"github.com/nikiforosFreespirit/msdb5/app/action/nextphase"
	"github.com/nikiforosFreespirit/msdb5/app/action/nextplayer"
	"github.com/nikiforosFreespirit/msdb5/app/action/phasesupplier"
	"github.com/nikiforosFreespirit/msdb5/app/game"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

// Orchestrator struct
type Orchestrator struct {
	game *game.Game
}

// NewAction func
func NewAction(side bool) app.Action {
	o := new(Orchestrator)
	o.game = game.NewGame(side)
	return o
}

// Action func
func (o *Orchestrator) Action(request, origin string) (all, me string, err error) {
	data := strings.Split(request, "#")
	currentPlayer := o.game.PlayerInTurn()
	// phase step
	currentPhase := o.game.CurrentPhase()
	inputPhase := phasesupplier.InputAction(data[0]).Phase()
	if currentPhase != inputPhase {
		return "", "", fmt.Errorf("Phase is not %d but %d", inputPhase, currentPhase)
	}
	// find step
	finder := NewFinder(data[0], request, origin, currentPlayer)
	_, p, err := o.game.Players().Find(finder.Find)
	if err != nil {
		return
	}
	// do step
	actionExec := NewExecuter(data[0], request, origin, o)
	if err = actionExec.Do(p); err != nil {
		return
	}

	// log action to file
	toFile(actionExec, p, o.game)

	// next player step
	nextPlayer := nextplayer.NewPlayerChanger(currentPhase, o.game.Players(), o.game.PlayedCards(), o.game.BriscolaSeed())
	o.game.NextPlayer(nextPlayer.NextPlayer)

	// next phase
	isSideDeckUsed := len(*o.game.SideDeck()) > 0
	nextPhase := nextphase.NewChanger(currentPhase, o.game.Players(), isSideDeckUsed, request)
	o.game.NextPhase(nextPhase.NextPhase())

	// log action to players
	all = infoForAll(currentPhase, *o.game)
	me = infoForMe(*currentPlayer, currentPhase, *o.game)
	o.game.Log(request, origin, err)

	// clean phase
	cleaner := clean.NewCleaner(o.game.PlayedCards())
	if cleaner != nil && len(*o.game.PlayedCards()) >= 5 {
		cleaner.Clean()
	}

	// process end game
	phaseAtEndTurn := o.game.CurrentPhase()
	if phaseAtEndTurn == game.End {
		scorers := make([]player.Scorer, 0)
		for _, p := range o.game.Players() {
			scorers = append(scorers, p)
		}
		scoreTeam1, scoreTeam2 := team.Score(p, o.game.Companion(), scorers...)
		all, me, err = fmt.Sprintf("Callers: %+v; Others: %+v", scoreTeam1, scoreTeam2), "", nil
	}
	return
}
