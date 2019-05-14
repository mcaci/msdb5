package orchestrator

import (
	"fmt"
	"strings"

	"github.com/nikiforosFreespirit/msdb5/app"
	"github.com/nikiforosFreespirit/msdb5/app/action"
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
	currentPhase := o.game.CurrentPhase()
	inputAction := phasesupplier.InputAction(data[0]).Phase()
	// phase step
	err = phaseStep(inputAction, currentPhase)
	if err != nil {
		return
	}
	// find step
	finder := NewFinder(data[0], request, origin, currentPlayer)
	p, err := findStep(finder, o.game.Players())
	if err != nil {
		return
	}
	// do step
	actionExec := NewExecuter(data[0], request, origin, o)
	err = playStep(actionExec, p)
	if err != nil {
		return
	}

	// log action to file
	toFile(actionExec, p, o.game)

	// next player step
	nextPlayer := nextplayer.NewPlayerChanger(inputAction, o.game.Players(), o.game.Board().PlayedCards(), o.game.BriscolaSeed())
	nextPlayerStep(nextPlayer, o.game)
	// next phase
	isSideDeckUsed := len(*o.game.Board().SideDeck()) > 0
	nextPhase := nextphase.NewChanger(inputAction, o.game.Players(), isSideDeckUsed, request)
	nextPhaseStep(nextPhase, o.game)

	all = infoForAll(currentPhase, *o.game)
	me = infoForMe(*currentPlayer, currentPhase, *o.game)
	o.game.Log(request, origin, err)

	// clean phase
	cleaner := clean.NewCleaner(o.game.Board().PlayedCards())
	if cleaner != nil && len(*o.game.Board().PlayedCards()) >= 5 {
		cleaner.Clean()
	}

	phaseAtEndTurn := o.game.CurrentPhase()
	if phaseAtEndTurn == game.End {
		all, me, err = endGame(o.game.Players(), o.game.Companion())
	}
	return
}

func phaseStep(input, current game.Phase) (err error) {
	if input != current {
		err = fmt.Errorf("Phase is not %d but %d", input, current)
	}
	return
}
func findStep(finder action.Finder, players team.Players) (*player.Player, error) {
	_, p, err := players.Find(finder.Find)
	return p, err
}
func playStep(executer action.Executer, p *player.Player) error {
	return executer.Do(p)
}
func nextPlayerStep(next action.NextPlayerSelector, g *game.Game) {
	g.NextPlayer(next.NextPlayer)
}
func nextPhaseStep(nextSel action.NextPhaseChanger, g *game.Game) {
	g.NextPhase(nextSel.NextPhase())
}
