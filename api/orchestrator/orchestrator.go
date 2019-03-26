package orchestrator

import (
	"fmt"
	"strings"

	"github.com/nikiforosFreespirit/msdb5/api"
	"github.com/nikiforosFreespirit/msdb5/api/action"
	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
	"github.com/nikiforosFreespirit/msdb5/team"
)

// Orchestrator struct
type Orchestrator struct {
	game *game.Game
}

// NewAction func
func NewAction(side bool) api.Action {
	o := new(Orchestrator)
	o.game = game.NewGame(side)
	return o
}

// Action func
func (o *Orchestrator) Action(request, origin string) (all, me string, err error) {
	data := strings.Split(request, "#")
	currentPlayer := o.game.PlayerInTurn()
	inputAction := action.InputAction(data[0])
	// phase step
	err = phaseStep(inputAction, o.game.CurrentPhase())
	if err != nil {
		return "", "", err
	}
	// find step
	var finder action.Finder
	switch data[0] {
	case "Join":
		finder = action.NewJoinFinder(request, origin)
	default:
		finder = action.NewPlayerFinder(origin, o.game.PlayerInTurn())
	}
	p, err := findStep(finder, o.game.Players())
	if err != nil {
		return "", "", err
	}
	// do step
	var actionExec action.Executer
	switch data[0] {
	case "Join":
		actionExec = action.NewJoin(request, origin)
	case "Auction":
		actionExec = action.NewAuction(request, origin, o.game.Players(), o.game.Board())
	case "Exchange":
		actionExec = action.NewExchangeCards(request, origin, o.game.Board().SideDeck())
	case "Companion":
		actionExec = action.NewCompanion(request, origin, o.game.Players(),
			o.game.SetCompanion)
	case "Card":
		actionExec = action.NewPlay(request, origin, o.game.Players(),
			o.game.Board().PlayedCards(), o.game.Board().SideDeck(), o.game.BriscolaSeed())
	}
	err = playStep(actionExec, p)
	if err != nil {
		return "", "", err
	}

	// next player
	var nextPlayer action.NextPlayerSelector
	switch data[0] {
	case "Join":
		nextPlayer = action.NewPlayerSelector()
	case "Auction":
		nextPlayer = action.NewAuction(request, origin, o.game.Players(), o.game.Board())
	case "Exchange":
		nextPlayer = action.NewExchangeCards(request, origin, o.game.Board().SideDeck())
	case "Companion":
		nextPlayer = action.NewCompanion(request, origin, o.game.Players(),
			o.game.SetCompanion)
	case "Card":
		nextPlayer = action.NewPlay(request, origin, o.game.Players(),
			o.game.Board().PlayedCards(), o.game.Board().SideDeck(), o.game.BriscolaSeed())
	}
	nextPlayerStep(nextPlayer, o.game)

	// next phase
	var nextPhase action.NextPhaseChanger
	switch data[0] {
	case "Join":
		nextPhase = action.NewPhaseChanger(o.game.Players())
	case "Auction":
		nextPhase = action.NewAuction(request, origin, o.game.Players(), o.game.Board())
	case "Exchange":
		nextPhase = action.NewExchangeCards(request, origin, o.game.Board().SideDeck())
	case "Companion":
		nextPhase = action.NewCompanion(request, origin, o.game.Players(),
			o.game.SetCompanion)
	case "Card":
		nextPhase = action.NewPlay(request, origin, o.game.Players(),
			o.game.Board().PlayedCards(), o.game.Board().SideDeck(), o.game.BriscolaSeed())
	}
	nextPhaseStep(nextPhase, o.game)

	all, me = fmt.Sprintf("Game: %+v", *o.game), fmt.Sprintf("%+v", currentPlayer)
	o.game.Log(request, origin, err)
	if o.game.CurrentPhase() == game.End {
		all, me, err = endGame(o.game.Players(), o.game.Companion())
	}
	return
}

func phaseStep(current action.PhaseSupplier, gamePhase game.Phase) (err error) {
	if gamePhase != current.Phase() {
		err = fmt.Errorf("Phase is not %d but %d", gamePhase, current.Phase())
	}
	return
}
func findStep(finder action.Finder, players playerset.Players) (*player.Player, error) {
	return players.Find(finder.Find)
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
